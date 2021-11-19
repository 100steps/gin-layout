package dao

import (
	"github.com/100steps/gin-layout/data/mysql"
	"github.com/100steps/gin-layout/util/paginator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `gorm:"type:varchar(12);not null;unique"`
	Password string `gorm:"type:varchar(64);not null;"`
	NickName string `gorm:"type:varchar(32);not null"`        // 昵称
	Sex      int    `gorm:"type:smallint;not null;default:0"` // 性别
}

// 创建一个用户
func (this *User) Create() error {
	if err := mysql.DB.Create(this).Error; err != nil {
		return err
	}
	return nil
}

// 根据openid判断用户是否存在
func (this *User) Exists() (bool, error) {
	var num int64
	if err := mysql.DB.Model(this).Where("account = ?", this.Account).Count(&num).Error; err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	return true, nil
}

// 根据account获取用户数据注入到结构体
func (this *User) FindByAccount(account string) error {
	return mysql.DB.Where("account = ?", account).Take(this).Error
}

// 根据id获取用户数据注入到结构体
func (this *User) FindById(id uint) error {
	return mysql.DB.Where("id = ?", id).Take(this).Error
}

func (this *User) GetAll() ([]User, error) {
	var res []User
	err := mysql.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *User) GetUsersPage(page, pageSize int) ([]User, error) {
	var res []User
	err := mysql.DB.Scopes(paginator.Paginate(page, pageSize)).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
