package utils

import (
	"mime/multipart"
	"os"
	"tes-face/config"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader, typeName string) (string, string, error)
	UploadFileByFile(file *os.File, typeName string, path string) (string, string, error)
	DeleteFile() error
}

func NewOss() OSS {
	switch config.GLOBAL.Type {
	case "local":
		return &Local{}
	case "tencent-cos":
		return &Local{}
	default:
		return &Local{}
	}
}
