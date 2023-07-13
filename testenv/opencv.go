package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())

	window := gocv.NewWindow("Hello")

	img := gocv.IMRead("bona.jpg", gocv.IMReadColor)

	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", "bona.jpg")
		return
	}

	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
