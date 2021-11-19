package migration

import (
	"log"

	"github.com/100steps/gin-layout/data/mysql"
	"gorm.io/gorm"
)

// 版本管理的计数
var migrationVer int

type Migration struct {
	gorm.Model
}

// 执行迁移: 这个方法接收一个包含具体迁移命令的函数，并自动进行版本管理
func migrate(f func() error) {
	migrationVer++
	if foundMigrationById(migrationVer) {
		return
	}
	if err := f(); err != nil {
		panic(err)
	}
	createMigration()
	log.Printf("Migration %d finished.\n", migrationVer)
}

// 根据id判断版本是否存在
func foundMigrationById(id int) bool {
	var count int64
	mysql.DB.Model(&Migration{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return true
}

// 创建迁移的版本
func createMigration() error {
	if err := mysql.DB.Create(&Migration{}).Error; err != nil {
		return err
	}
	return nil
}
