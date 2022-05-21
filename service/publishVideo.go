package service

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	. "github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
)

var c *gin.Context

type PublishFlow struct {
	user      *User
	video     *Video
	FinalName string
}

func NewPublishFlowInstance(ctx *gin.Context) *PublishFlow {
	c = ctx
	return &PublishFlow{}
}

//先检查用户是否存在
func (p *PublishFlow) getUser() error {
	token := c.PostForm("token")
	u, err := GetUser_token(token)
	if err != nil {
		return err
	}
	p.user = u
	return nil
}

func (p *PublishFlow) saveFile() error {
	data, err := c.FormFile("data")
	if err != nil {
		return err
	}
	filename := filepath.Base(data.Filename)
	videoId := VideoId()
	//这里使用 用户id + 视频Id + 视频名称做存储
	finalName := fmt.Sprintf("%d-%d_%s", p.user.Id, videoId, filename)
	p.FinalName = finalName
	save := filepath.Join("./public/", finalName)
	//本地地址
	playPath := fmt.Sprintf("http://192.168.212.205:8080/static/%s", finalName)
	err = c.SaveUploadedFile(data, save)
	if err != nil {
		return err
	}
	p.video = &Video{
		Id:      videoId,
		Author:  *p.user,
		PlayUrl: playPath,
		Created: time.Now().Unix(),
	}
	err = SaveVideo(p.video)
	if err != nil {
		return err
	}
	return nil
}

func (p *PublishFlow) Do() error {
	err := p.getUser()
	if err != nil {
		return errors.New("user doesn't exist")
	}
	err = p.saveFile()
	if err != nil {
		return err
	}
	return nil
}
