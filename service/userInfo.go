package service

import . "github.com/RaymondCode/simple-demo/repository"

func UserInfo(token string) (*User, error) {
	return GetUser_token(token)
}
