package utils

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"tes-face/config"
	"time"
)

type Local struct{}

func (l *Local) UploadFileByFile(file *os.File, typeName string, filePath string) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Name())
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Name(), ext)
	name = MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(config.FILEPATH.GeneratePath+"/"+typeName, os.ModePerm)
	if mkdirErr != nil {
		config.TES_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := config.FILEPATH.GeneratePath + "/" + typeName + "/" + filename
	out, createErr := os.Create(p)
	if createErr != nil {
		config.TES_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	open, openError := os.Open(filePath)
	// 创建文件 defer 关闭
	defer open.Close()
	if openError != nil {
		config.TES_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	// 创建文件 defer 关闭
	defer out.Close()
	// 传输（拷贝）文件
	_, copyErr := io.Copy(out, open)
	if copyErr != nil {
		config.TES_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

func (*Local) UploadFile(file *multipart.FileHeader, typeName string) (string, string, error) {
	name := file.Filename
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(config.FILEPATH.Path+"/"+typeName, os.ModePerm)
	if mkdirErr != nil {
		config.TES_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := config.FILEPATH.Path + "/" + typeName + "/" + name
	f, openError := file.Open() // 读取文件
	if openError != nil {
		config.TES_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		config.TES_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭
	// 传输（拷贝）文件
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		config.TES_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, name, nil
}

func (*Local) DeleteFile() error {
	p := config.FILEPATH.GeneratePath
	if err := os.RemoveAll(p); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	g := config.FILEPATH.Path
	if err := os.RemoveAll(g); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}
