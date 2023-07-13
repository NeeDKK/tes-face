package handleface

import (
	"errors"
	"fmt"
	face "github.com/Kagami/go-face"
	"github.com/nfnt/resize"
	"go.uber.org/zap"
	"gocv.io/x/gocv"
	"image"
	"image/draw"
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
	// 初始化识别器
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()
	fmt.Println("Recognizer Initialized")
	// 调用该方法，传入路径。返回面部数量和任何错误
	faces, err := rec.RecognizeFile(imageFilePath)
	if err != nil {
		fmt.Println("无法识别: %v", err)
		return "", "", errors.New("picture recoginzation fail, err:" + err.Error())
	}
	// 打印人脸数量
	fmt.Println("图片人脸数量: ", len(faces))
	if len(faces) == 0 {
		return imageFilePath, origin, nil
	}
	for _, f := range faces {
		//马赛克图片
		//open, err := os.Open(imageFilePath)
		//decode, _ := jpeg.Decode(open)
		//process, err := stackblur.Process(decode, 100)
		//if err != nil {
		//	return "", "", err
		//}
		//create, err := os.Create(masakePath)
		//jpeg.Encode(create, process, nil)
		//in, _ := os.Open(masakePath)
		//out, _ := os.Create("./out.jpg")
		//err = Clip(in, out, 0, 0, f.Rectangle.Max.X, f.Rectangle.Max.Y, f.Rectangle.Min.X, f.Rectangle.Min.Y, 100)
		//
		//imgb, _ := os.Open("./out.jpg")
		//img, _ := jpeg.Decode(imgb)

		//defer out.Close()
		//defer open.Close()
		//defer create.Close()
		//defer imgb.Close()
		//defer in.Close()

		imgb, _ := os.Open(masakePath)
		img, _ := png.Decode(imgb)
		defer imgb.Close()

		//目标图片
		wmb, _ := os.Open(imageFilePath)
		var watermark image.Image
		if strings.Contains(ext, "png") {
			watermark, _ = png.Decode(wmb)
		} else {
			watermark, _ = jpeg.Decode(wmb)
		}
		defer wmb.Close()
		b := watermark.Bounds()
		m := image.NewNRGBA(b)

		draw.Draw(m, b, watermark, image.ZP, draw.Src)
		draw.Draw(m, f.Rectangle, img, image.ZP, draw.Over)

		//new
		mkdirErr := os.MkdirAll(config.FILEPATH.GeneratePath+"/images/", os.ModePerm)
		if mkdirErr != nil {
			config.TES_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
			return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
		}
		imgw, _ := os.Create(path)
		jpeg.Encode(imgw, m, &jpeg.Options{100})
		defer imgw.Close()
		imageFilePath = path
	}
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
	mosaicImage := gocv.IMRead(mosaicImageFile, gocv.IMReadColor)
	if err != nil {
		fmt.Println(err)
		return videoFile, origin, err
	}
	defer mosaicImage.Close()

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
		for _, r := range rects {
			// 提取人脸区域
			faceRegion := frame.Region(r)

			// 调整马赛克图像大小为人脸区域大小
			resizedMosaic := gocv.NewMatWithSize(faceRegion.Rows(), faceRegion.Cols(), mosaicImage.Type())
			gocv.Resize(mosaicImage, &resizedMosaic, image.Point{X: faceRegion.Cols(), Y: faceRegion.Rows()}, 0, 0, gocv.InterpolationDefault)
			resizedMosaic.ConvertTo(&resizedMosaic, faceRegion.Type())
			// 将调整后的马赛克图像应用到人脸区域
			gocv.AddWeighted(faceRegion, 0.0, resizedMosaic, 1.0, 0.0, &faceRegion)

			// 释放资源
			resizedMosaic.Close()
			faceRegion.Close()
		}

		// 写入带有自定义马赛克的视频帧到输出文件
		output.Write(frame)

	}

	// 视频输出结束
	fmt.Println("视频输出完成")
	return outputFile, filename, nil
}
