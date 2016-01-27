package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/lazywei/go-opencv/opencv"
)

func main() {

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("cannot open camera")
	}
	defer cap.Release()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cascade := opencv.LoadHaarClassifierCascade(path.Join(cwd, "haarcascade_frontalface_alt.xml"))

	fmt.Println("Press ESC to quit")
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				faces := cascade.DetectObjects(img)
				for _, value := range faces {
					leftX := value.X()
					rightX := leftX + value.Width()
					lowY := value.Y()
					highY := lowY + value.Height()
					log.Println("Face Seen! LowLeft Corner: (", leftX, ", ", lowY, "), UpRight Corner: (", rightX, ", ", highY, ")")
				}

			} else {
				fmt.Println("nil image")
			}
		}
		key := opencv.WaitKey(1)

		if key == 27 {
			os.Exit(0)
		}
	}
}
