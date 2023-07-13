package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
	"tes-face/config"
	response "tes-face/entity"
	"tes-face/service"
)

type UploadController struct {
}

func UploadPic(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		config.TES_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file := service.UploadFile(header, "images")
	if err != nil {
		config.TES_LOG.Error("function file.Open() Filed", zap.Any("err", err.Error()))
		response.FailWithMessage("license文件读取失败", c)
		return
	}
	fileNameReturn := url.QueryEscape(file.Name)
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+fileNameReturn)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+fileNameReturn)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(file.Key)
}

func UploadVideo(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		config.TES_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file := service.UploadFile(header, "video")
	if err != nil {
		config.TES_LOG.Error("function file.Open() Filed", zap.Any("err", err.Error()))
		response.FailWithMessage("license文件读取失败", c)
		return
	}
	fileNameReturn := url.QueryEscape(file.Name)
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+fileNameReturn)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+fileNameReturn)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(file.Key)
}

func ClearFile(c *gin.Context) {
	err := service.DeleteFile()
	if err != nil {
		config.TES_LOG.Error("delete file Filed", zap.Any("err", err.Error()))
		return
	}
	response.OkWithData("删除成功", c)
	return
}
