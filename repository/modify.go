package repository

import (
	"errors"
)

func SaveVideo(v *Video) error {
	return db.Create(v).Error
}

func AddUser(user *User, token string) error {
	has := int64(0)
	s := db.Model(&User{}).Where("Name = ?", user.Name).Find(&User{})
	s.Count(&has)
	if has > 0 {
		return errors.New("userName already exist")
	}
	user.Id = UserId()
	result := db.Create(&StoredUser{
		Token: token,
		User:  *user,
	})
	return result.Error
}
