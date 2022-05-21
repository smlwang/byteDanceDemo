package repository

import (
	"errors"

	"gorm.io/gorm"
)

//普通视频流推送
func DefalutVideo(latest_time int64) (*[]Video, int64, error) {
	res := []Video{}
	result := db.Debug().Model(&Video{}).Order("Created desc").Where("Created <= ?", latest_time).Preload("Author").Limit(30).Find(&res)
	if len(res) == 0 {
		return nil, -1, errors.New("no more video")
	}
	sv := Video{}
	result.Last(&sv)
	return &res, sv.Created, result.Error
}

//获取用户发布视频
func UserVideo(token string) (*[]Video, error) {
	u := StoredUser{}
	result := db.Model(&StoredUser{}).Where("Token = ?", token).First(&u)
	if u.Token == "" || result.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user not exist")
	}
	v := []Video{}
	result = db.Where(&Video{Author: User{Id: u.UserKey_Stored}}).Preload("Author").Find(&v)
	return &v, result.Error
}

func GetUser_token(token string) (*User, error) {
	u := StoredUser{}
	result := db.Model(&StoredUser{}).Where("Token = ?", token).Preload("User").First(&u)
	if u.Token == "" || result.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("userName or passWord wrong")
	}
	return &u.User, nil
}
func GetIdInfo() IdInfo {
	i := IdInfo{}
	db.Model(&IdInfo{}).First(&i)
	return i
}
