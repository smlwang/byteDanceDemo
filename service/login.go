package service

import (
	. "github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/tool"
)

type LoginFlow struct {
	User   *User
	Name   string
	Passwd string
	Token  string
	Err    error
}

func NewLoginFlowInstance(name string, passwd string) *LoginFlow {
	return &LoginFlow{
		Name:   name,
		Passwd: passwd,
	}
}
func (l *LoginFlow) Do() {
	l.Token = tool.Token(l.Name + l.Passwd)
	l.User, l.Err = GetUser_token(l.Token)
}
