package service

import (
	. "github.com/RaymondCode/simple-demo/repository"
)

func GetList(token string) (*[]Video, error) {
	return UserVideo(token)
}
