package service

import (
	"mime/multipart"
	"strings"
	"tes-face/entity"
	"tes-face/handleface"
	upload "tes-face/utils"
)

func UploadFile(header *multipart.FileHeader, typeName string) (err error, file entity.ExaFileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header, typeName)
	if uploadErr != nil {
		panic(err)
	}
	s := strings.Split(header.Filename, ".")
	f := entity.ExaFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	if typeName == "images" {
		path, name, err := handleface.HandlePic(f.Url, f.Name)
		if err != nil {
			return err, f
		}
		f.Key = path
		f.Name = name
	}
	if typeName == "video" {
		path, name, err := handleface.HandleVideo(f.Url, f.Name)
		if err != nil {
			return err, f
		}
		f.Key = path
		f.Name = name
	}
	return err, f
}

func DeleteFile() error {
	oss := upload.NewOss()
	return oss.DeleteFile()
}
