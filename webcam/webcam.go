package webcam

import (
	"gocv.io/x/gocv"
)

func CaptureWebcam() {
	webcam, _ := gocv.OpenVideoCapture(2)
	webcam2, _ := gocv.OpenVideoCapture(4)
	window := gocv.NewWindow("Hello")
	window2 := gocv.NewWindow("Hello2")
	img := gocv.NewMat()
	img2 := gocv.NewMat()
	for {

		webcam.Read(&img)
		webcam2.Read(&img2)
		window.IMShow(img)
		window2.IMShow(img2)
		window.WaitKey(1)
		window2.WaitKey(1)
	}
}
