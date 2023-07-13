package main

import (
	"fmt"

	"github.com/Kagami/go-face"
)

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	modelDir  = "models"
	imagesDir = "images"
)

func main() {
	fmt.Println("Face Recognition...")

	// 初始化识别器
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")
}
