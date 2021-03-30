package libdetector

import "gocv.io/x/gocv"

type Detector interface {
	Detect(img gocv.Mat, threshold float32) Boxes
	DetectBlob(blob gocv.Mat, threshold float32) Boxes
}

type Box struct {
	class int
	rect  struct {
		min [2]float32
		max [2]float32
	}
	conf float32
}

type Boxes []Box
