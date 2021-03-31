package main

import (
	"fmt"
	"github.com/umbrella-evgeny-nefedkin/goreco"
	"gocv.io/x/gocv"
)

const NN_MODEL_PATH = "example/model/ssd.pb"
const NN_MODEL_CONFIG = "example/model/ssd.pbtxt"

func main() {

	deviceID := 0 // local cam

	// open capture device
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	// Create pre-config
	Config := goreco.DefaultSSDConfig

	// Set model path
	Config.ModelPath = NN_MODEL_PATH
	Config.ConfigPath = NN_MODEL_CONFIG

	// Create detector
	Detector := goreco.NewDetector(Config)

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// Call detection
		result := Detector.Detect(img, 0.6, true)

		// Mark detected regions on original frame
		goreco.MarkObjects(&img, result, true)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
