package migration

import (
	dao "github.com/100steps/gin-layout/dao"
	"github.com/100steps/gin-layout/data/mysql"
)

func init() {
	if err := mysql.DB.AutoMigrate(&Migration{}); err != nil {
		panic(err)
	}
}

func init() {
	migrate(func() error {
		return mysql.DB.AutoMigrate(&dao.User{})
	})
}

//NOTE: 后续版本一直在下面平铺下去
//NOTE: 这里的代码只能增加，绝对不能删除或者改动！！！不然会有严重后果！！
