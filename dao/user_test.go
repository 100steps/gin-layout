package dao

import (
	"testing"

	"github.com/100steps/gin-layout/data/mysql"
)

func TestCreate(t *testing.T) {
	user := &User{
		Account:  "ForSeason",
		Password: "smjb",
		NickName: "ForSeason",
		Sex:      0,
	}
	if err := mysql.DB.Create(user).Error; err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}
