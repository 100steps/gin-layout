package middleware

import (
	"github.com/100steps/gin-layout/dao"
	"github.com/100steps/gin-layout/util/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 判断用户是否登陆，登陆了就获取session里面的用户数据并写入context，
// 没登录就直接停止后续逻辑执行
func (this *Middleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取session的user
		// TODO: 目前的逻辑是获取缓存的id之后去db查一次，性能有点难以接受，需要优化
		session := sessions.Default(ctx)
		sessionUser := session.Get("user")
		if sessionUser == nil {
			ctx.JSON(e.UNAUTHORIZED, e.ErrMsgResponse{Message: e.GetMsg(e.ERROR_AUTH)})
			ctx.Abort()
			return
		}
		user, err := this.userService.GetUserById(sessionUser.(*dao.User).ID)
		if err != nil {
			ctx.JSON(e.UNAUTHORIZED, e.ErrMsgResponse{Message: e.GetMsg(e.ERROR_AUTH)})
			ctx.Abort()
			return
		}
		ctx.Keys["user"] = user
		ctx.Next()
	}
}
