package service

import (
	"github.com/100steps/gin-layout/dao"
)

var userDao *dao.User = &dao.User{}

type UserService struct {
	userDao *dao.User
}

func NewUserService() *UserService {
	return &UserService{userDao: &dao.User{}}
}

// 如果用户不存在，就创建一个用户
func (this *UserService) CreateUserIfNotExists(user *dao.User) error {
	exists, err := user.Exists()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return user.Create()
}

func (this *UserService) CreateUser(account, password, nickname string, sex int) error {
	user := &dao.User{
		Account:  account,
		Password: password,
		NickName: nickname,
		Sex:      sex,
	}
	return user.Create()
}

func (this *UserService) GetUserById(id uint) (*dao.User, error) {
	user := &dao.User{}
	if err := user.FindById(id); err != nil {
		return nil, err
	}
	return user, nil
}

func (this *UserService) GetUserByAccount(account string) (*dao.User, error) {
	user := &dao.User{}
	if err := user.FindByAccount(account); err != nil {
		return nil, err
	}
	return user, nil
}

func (this *UserService) GetAllUsers() ([]dao.User, error) {
	return userDao.GetAll()
}

func (this *UserService) GetUsersPage(page, pageSize int) ([]dao.User, error) {
	return userDao.GetUsersPage(page, pageSize)
}
