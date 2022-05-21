package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:wang200315@tcp(localhost:3306)/DB1?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB

type StoredUser struct {
	Token          string
	UserKey_Stored int64
	User           User `gorm:"foreignKey:UserKey_Stored"`
}

func Init() {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = database
	err = db.AutoMigrate(&StoredUser{}, &Video{}, &IdInfo{})
	if err != nil {
		log.Fatal(err)
	}
	err = idInit()
	if err != nil {
		log.Fatal(err)
	}
}

//测试用, 后续应该更改方案
func idInit() error {
	i := IdInfo{}
	result := db.Model(&IdInfo{}).First(&i)
	userId = i.UserId_Max
	videoId = i.VideoId_Max
	commentId = i.CommentId_Max
	if result.Error == gorm.ErrRecordNotFound {
		db.Create(&IdInfo{
			UserId_Max:    0,
			VideoId_Max:   0,
			CommentId_Max: 0,
		})
		return nil
	}
	return result.Error
}

//将id状态保留, 防止id重复
func End() {
	i := IdInfo{}
	db.First(&i)
	i.UserId_Max = userId
	i.VideoId_Max = videoId
	i.CommentId_Max = commentId
	db.Save(&i)
}
