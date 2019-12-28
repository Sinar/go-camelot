package camelot

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
	"os"
)

// Here is the gocv method

// Example from: https://raw.githubusercontent.com/hybridgroup/gocv/master/cmd/facedetect/main.go
// What it does:
//
// This example uses the CascadeClassifier class to detect faces,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// How to run:
//
// facedetect [camera ID] [classifier XML file]
//
// 		go run ./cmd/facedetect/main.go 0 data/haarcascade_frontalface_default.xml
//
// +build example

func runExample() {
	filename := os.Args[1]

	mat := gocv.IMRead(filename, gocv.IMReadColor)

	matCanny := gocv.NewMat()
	matLines := gocv.NewMat()

	window := gocv.NewWindow("detected lines")

	gocv.Canny(mat, &matCanny, 50, 200)
	gocv.HoughLinesP(matCanny, &matLines, 1, math.Pi/180, 80)

	fmt.Println(matLines.Cols())
	fmt.Println(matLines.Rows())
	for i := 0; i < matLines.Rows(); i++ {

		fmt.Println("ROW: ", i)
		// spew.Dump(matLines.GetVeciAt(i, 0))

		pt1 := image.Pt(int(matLines.GetVeciAt(i, 0)[0]), int(matLines.GetVeciAt(i, 0)[1]))
		pt2 := image.Pt(int(matLines.GetVeciAt(i, 0)[2]), int(matLines.GetVeciAt(i, 0)[3]))
		spew.Dump(pt1)
		spew.Dump(pt2)
		gocv.Line(&mat, pt1, pt2, color.RGBA{0, 255, 0, 50}, 10)
	}

	for {
		window.IMShow(mat)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}

