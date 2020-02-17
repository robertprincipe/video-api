package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertprincipe/video/controller"
	"github.com/robertprincipe/video/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, videoController.Save(ctx))
	})

	server.Run(":9000")
}
