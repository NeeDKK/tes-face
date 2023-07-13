package main

import (
	"github.com/gin-gonic/gin"
	"tes-face/config"
	"tes-face/controller"
	"tes-face/entity"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	//创建路由
	engine := gin.Default()
	engine.Use(config.Cors())
	engine.POST("/upload-pic", controller.UploadPic)
	engine.POST("/upload-video", controller.UploadVideo)
	engine.GET("/clear-file", controller.ClearFile)
	//启动gin服务
	engine.Run(":9999")
}

func init() {
	config.FILEPATH = &entity.FilePath{
		Path:         "./origin",
		GeneratePath: "./generate",
	}
	config.GLOBAL = &entity.OssInfo{
		Type: "local",
	}

}
