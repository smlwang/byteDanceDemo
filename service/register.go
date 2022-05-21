package service

import (
	"errors"

	. "github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/tool"
)

type RegistFlow struct {
	User   *User
	Passwd string
	Token  string
}

func NewRegistFlowInstance(user *User, passwd string) *RegistFlow {
	return &RegistFlow{
		User:   user,
		Passwd: passwd,
	}
}
func checkName(name string) error {
	if len(name) > 32 {
		return errors.New("name too long, can't excceed 32")
	}
	return nil
}
func checkPasswd(pwd string) error {
	l := len(pwd)
	if l < 6 {
		return errors.New("too short, least 6")
	}
	if l > 32 {
		return errors.New("password too long, can't excceed 32")
	}
	return nil
}
func (r *RegistFlow) Do() error {
	if err := checkName(r.User.Name); err != nil {
		return err
	}
	if err := checkPasswd(r.Passwd); err != nil {
		return err
	}
	token := tool.Token(r.User.Name + r.Passwd)
	r.Token = token
	return AddUser(r.User, token)
}
