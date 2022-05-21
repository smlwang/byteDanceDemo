package service

import (
	"time"

	"github.com/RaymondCode/simple-demo/tool"

	"github.com/RaymondCode/simple-demo/repository"
)

type GetVideoFlow struct {
	Latest_time int64
	Token       int64
	VideoList   *[]repository.Video
	NextTime    int64
}

func NewVideoFlowInstance(latest_time string, token string) *GetVideoFlow {
	return &GetVideoFlow{
		Latest_time: tool.Int64(latest_time, func() int64 {
			return time.Now().Unix()
		}),
		Token: tool.Int64(token, func() int64 {
			return 0
		}),
	}
}
func (v *GetVideoFlow) Do() error {
	videoList, nextTime, err := repository.DefalutVideo(v.Latest_time)
	if err != nil {
		return err
	}
	v.NextTime = nextTime
	v.VideoList = videoList
	return nil
}
