package repository

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDb(t *testing.T) {
	var dsn = "root:212180958zxc@tcp(localhost:3305)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = database

	sub := []DbFollow{}
	db.Model(&DbFollow{}).Where("id = ?", 7).Find(&sub)
	followeds := make([]DbUser, len(sub))
	for i, v := range sub {
		user := DbUser{}
		db.Model(&DbUser{}).Where("id = ?", v.FollowId).Find(&user)
		followeds[i] = user
	}
	fmt.Println(followeds)
}
