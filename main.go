package main

import (
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repository.Init()
	defer repository.End()
	initRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
