package main

import (
	"douyinProject/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repository.Init()
	initRouter(r)

	r.Run(":8001") // listen and serve on 0.0.0.0:8001 (for windows "localhost:8001")
}
