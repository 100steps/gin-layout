package controller

import (
	"encoding/gob"

	"github.com/100steps/gin-layout/dao"
	"github.com/100steps/gin-layout/service"
	"github.com/100steps/gin-layout/util/e"
	"github.com/100steps/gin-layout/util/encryptor"
	"github.com/100steps/gin-layout/util/paginator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	gob.Register(&dao.User{})
}

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// 获取用户的列表，带分页
func (this *UserController) Index(ctx *gin.Context) {
	// userList, err := this.userService.GetAllUsers()
	page, pageSize := paginator.GetPageParams(ctx)
	userList, err := this.userService.GetUsersPage(page, pageSize)
	if err != nil {
		ctx.JSON(e.ERROR, &e.ErrMsgResponse{Message: err.Error()})
		return
	}
	ctx.JSON(200, userList)
}

// 获取已登录用户的信息
func (this *UserController) ShowMe(ctx *gin.Context) {
	ctx.JSON(200, ctx.Keys["user"].(*dao.User))
}

type UserLoginReq struct {
	Account  string `binding:"required"`
	Password string `binding:"required"`
}

func (this *UserController) Login(ctx *gin.Context) {
	var req UserLoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(e.INVALID_PARAMS, &e.ErrMsgResponse{Message: e.GetMsg(e.INVALID_PARAMS)})
		return
	}
	user, err := this.userService.GetUserByAccount(req.Account)
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			ctx.JSON(e.ERROR_AUTH, &e.ErrMsgResponse{Message: e.GetMsg(e.ERROR_AUTH)})
		} else {
			ctx.JSON(e.ERROR, &e.ErrMsgResponse{Message: err.Error()})
		}
		return
	}
	req.Password = encryptor.MD5(req.Password)
	if user.Password != req.Password {
		ctx.JSON(e.ERROR_AUTH, &e.ErrMsgResponse{Message: e.GetMsg(e.ERROR_AUTH)})
		return
	}
	if err := this.login(ctx, *user); err != nil {
		ctx.JSON(e.ERROR, &e.ErrMsgResponse{Message: err.Error()})
		return
	}
	user.Password = ""
	ctx.JSON(200, user)
}

type UserRegisterReq struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
	Nickname string `form:"nickname" binding:"required"`
	Sex      int    `form:"sex"`
}

func (this *UserController) Register(ctx *gin.Context) {
	var req UserRegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(e.INVALID_PARAMS, &e.ErrMsgResponse{Message: err.Error()})
		return
	}
	req.Password = encryptor.MD5(req.Password)
	if err := this.userService.CreateUser(req.Account, req.Password, req.Nickname, req.Sex); err != nil {
		ctx.JSON(e.ERROR, &e.ErrMsgResponse{Message: e.GetMsg(e.ERROR_USER_CREATE_FAIL)})
		return
	}
	ctx.JSON(200, req)
}

// 通用登陆逻辑：
// 1. 首先判断该用户是否存在，不存在则创建一个用户
// 2. 将用户数据写入session
// 3. 向context写入请求结果
// NOTE: session的数据和数据库里面的数据不一定同步，请不要在业务中直接使用sessino里面的user数据！
func (this *UserController) login(ctx *gin.Context, user dao.User) error {
	session := sessions.Default(ctx)
	session.Set("user", user)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
