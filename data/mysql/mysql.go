package mysql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/forseason/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	user := env.Get("MYSQL_USER", "forseason")
	password := env.Get("MYSQL_PASSWORD", "root")
	host := env.Get("MYSQL_HOST", "127.0.0.1:3306")
	name := env.Get("MYSQL_NAME", "demo")
	maxIdle, err := strconv.Atoi(env.Get("MYSQL_MAX_IDLE", "10"))
	if err != nil {
		panic(err)
	}
	maxActive, err := strconv.Atoi(env.Get("MYSQL_MAX_ACTIVE", "50"))
	if err != nil {
		panic(err)
	}
	maxLifetime, err := strconv.Atoi(env.Get("MYSQL_MAX_LIFETIME", "3600"))
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, name)),
		&gorm.Config{
			// 启用这个选项会把默认的事务管理关闭，不追求性能的话可以把这行注释掉
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxActive)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
}
