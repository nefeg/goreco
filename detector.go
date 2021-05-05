package goreco

import "gocv.io/x/gocv"

type Detector interface {
	Detect(img gocv.Mat, threshold float32, all bool) []Box
	DetectBlob(blob gocv.Mat, threshold float32, all bool) []Box
}

type Box struct {
	class int
	rect  struct {
		min [2]float32
		max [2]float32
	}
	conf float32
}
