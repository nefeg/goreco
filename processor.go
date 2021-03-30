package libdetector

import "gocv.io/x/gocv"

type Processor interface {
	Process(detectionResult gocv.Mat, threshold float32) Boxes
}
