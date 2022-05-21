package controller

import (
	"log"
	"net/http"
	"time"

	. "github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func bad(err error, c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		VideoList: nil,
		NextTime:  time.Now().Unix(),
	})
	log.Println(err)
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latest_time := c.Query("latest_time")
	token := c.Query("token")
	workFlow := service.NewVideoFlowInstance(latest_time, token)
	err := workFlow.Do()
	if err != nil {
		bad(err, c)
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: *(workFlow.VideoList),
		NextTime:  workFlow.NextTime,
	})
}
