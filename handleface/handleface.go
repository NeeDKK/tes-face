package handleface

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"go.uber.org/zap"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"tes-face/config"
	"time"
)

const (
	masakePath = "masaike.png"
	modelDir   = "models"
)

func HandlePic(imageFilePath, origin string) (string, string, error) {
	fmt.Println("Face Recognition...")
	// 读取文件后缀
	ext := path.Ext(origin)
	// 读取文件名并加密
	name := strings.TrimSuffix(origin, ext)
	name = name + "-handle"
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	path := config.FILEPATH.GeneratePath + "/images/" + filename
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file")
		return "", "", errors.New("读取模型文件失败")
	}
	fmt.Println("Recognizer Initialized")
	// Read the image you want to analyze.
	img := gocv.IMRead(imageFilePath, gocv.IMReadColor)

	if img.Empty() {
		fmt.Println("读取图片失败")
		return "", "", errors.New("读取图片失败")
	}
	defer img.Close()
	// Detect faces in the image.
	rects := classifier.DetectMultiScale(img)
	fmt.Printf("识别特征人脸 %d \n", len(rects))
	if len(rects) == 0 {
		return imageFilePath, origin, nil
	}
	// 加载你想要覆盖的图片并将其转换为gocv.Mat格式
	overlayImg := gocv.IMRead(masakePath, gocv.IMReadColor)
	for _, f := range rects {
		// 确定将图片覆盖的位置（这里示例为在人脸的左上角）
		posX := f.Min.X
		posY := f.Min.Y
		overlayWidth := f.Max.X - f.Min.X
		overlayHeight := f.Max.Y - f.Min.Y

		// 确保裁剪区域不超过overlayImg的大小
		if overlayWidth > overlayImg.Cols() {
			overlayWidth = overlayImg.Cols()
		}
		if overlayHeight > overlayImg.Rows() {
			overlayHeight = overlayImg.Rows()
		}

		// 获取裁剪后的覆盖图片
		overlayCropped := overlayImg.Region(image.Rect(0, 0, overlayWidth, overlayHeight))

		// 获取原始图像中指定位置的区域
		targetRegion := img.Region(image.Rect(posX, posY, posX+overlayWidth, posY+overlayHeight))

		// 将裁剪后的覆盖图片覆盖在原始图像的指定位置
		overlayCropped.CopyTo(&targetRegion)

		// 绘制人脸边界框
		gocv.Rectangle(&img, f, color.RGBA{0, 255, 0, 0}, 2)
	}
	//new
	mkdirErr := os.MkdirAll(config.FILEPATH.GeneratePath+"/images/", os.ModePerm)
	if mkdirErr != nil {
		config.TES_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	gocv.IMWrite(path, img)
	return path, filename, nil
}

func Clip(in io.Reader, out io.Writer, wi, hi, x0, y0, x1, y1, quality int) (err error) {
	err = errors.New("unknow error")
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	var origin image.Image
	var fm string
	origin, fm, err = image.Decode(in)
	if err != nil {
		log.Println(err)
		return err
	}

	if wi == 0 || hi == 0 {
		wi = origin.Bounds().Max.X
		hi = origin.Bounds().Max.Y
	}
	var canvas image.Image
	if wi != origin.Bounds().Max.X {
		//先缩略
		canvas = resize.Thumbnail(uint(wi), uint(hi), origin, resize.Lanczos3)
	} else {
		canvas = origin
	}

	switch fm {
	case "jpeg":
		img := canvas.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
	case "png":
		switch canvas.(type) {
		case *image.NRGBA:
			img := canvas.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := canvas.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		}
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}

func HandleVideo(videoFile, origin string) (string, string, error) {
	// 设置使用的CPU核心数
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Video Face Recognition...")
	// 读取文件后缀
	ext := path.Ext(origin)
	// 读取文件名并加密
	name := strings.TrimSuffix(origin, ext)
	name = name + "-handle"
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	outputFile := config.FILEPATH.GeneratePath + "/videos/" + filename
	mkdirErr := os.MkdirAll(config.FILEPATH.GeneratePath+"/videos/", os.ModePerm)
	if mkdirErr != nil {
		config.TES_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 打开视频文件
	video, err := gocv.VideoCaptureFile(videoFile)
	if err != nil {
		fmt.Println(err)
		return videoFile, origin, err
	}
	defer video.Close()

	// 获取视频的属性
	width := int(video.Get(gocv.VideoCaptureFrameWidth))
	height := int(video.Get(gocv.VideoCaptureFrameHeight))
	fps := float64(video.Get(gocv.VideoCaptureFPS))

	// 创建视频编写器
	output, err := gocv.VideoWriterFile(outputFile, "avc1", fps, width, height, true)
	if err != nil {
		fmt.Println(err)
		return videoFile, origin, err
	}
	defer output.Close()
	// 加载人脸识别分类器
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("无法加载人脸分类器")
		return videoFile, origin, err
	}

	// 加载自定义的马赛克图片
	mosaicImageFile := masakePath
	overlayImg := gocv.IMRead(mosaicImageFile, gocv.IMReadColor)
	if err != nil {
		fmt.Println(err)
		return videoFile, origin, err
	}
	defer overlayImg.Close()

	// Create channel to handle processed frames
	frameChan := make(chan gocv.Mat)

	// Start goroutine to process frames
	go processFrames(video, &classifier, overlayImg, frameChan)

	// Wait for the processing to finish and check for errors
	for frame := range frameChan {
		// Write processed frames to the output video
		output.Write(frame)
	}

	// 视频输出结束
	fmt.Println("视频输出完成")
	return outputFile, filename, nil
}

func processFrames(video *gocv.VideoCapture, classifier *gocv.CascadeClassifier, overlayImg gocv.Mat, frameChan chan<- gocv.Mat) {
	// 循环读取视频帧并进行处理
	frame := gocv.NewMat()
	defer frame.Close()
	for {
		if ok := video.Read(&frame); !ok {
			// 视频结束
			break
		}
		if frame.Empty() {
			continue
		}

		// 在当前帧中检测人脸
		rects := classifier.DetectMultiScale(frame)
		for _, f := range rects {
			// 提取人脸区域
			faceRegion := frame.Region(f)
			// 确定将图片覆盖的位置（这里示例为在人脸的左上角）
			posX := f.Min.X
			posY := f.Min.Y
			overlayWidth := f.Max.X - f.Min.X
			overlayHeight := f.Max.Y - f.Min.Y
			// 确保裁剪区域不超过overlayImg的大小
			if overlayWidth > overlayImg.Cols() {
				overlayWidth = overlayImg.Cols()
			}
			if overlayHeight > overlayImg.Rows() {
				overlayHeight = overlayImg.Rows()
			}
			// 获取裁剪后的覆盖图片
			overlayCropped := overlayImg.Region(image.Rect(0, 0, overlayWidth, overlayHeight))
			// 获取原始图像中指定位置的区域
			targetRegion := frame.Region(image.Rect(posX, posY, posX+overlayWidth, posY+overlayHeight))
			// 将裁剪后的覆盖图片覆盖在原始图像的指定位置
			overlayCropped.CopyTo(&targetRegion)
			// 绘制人脸边界框
			gocv.Rectangle(&frame, f, color.RGBA{0, 255, 0, 0}, 2)
			// 释放资源
			faceRegion.Close()
		}
		frameChan <- frame
	}

	// Close the frameChan channel when processing is complete
	close(frameChan)
}
