package goreco

import (
	"gocv.io/x/gocv"
	"image"
	"log"
	"os"
)

type detector struct {
	net    *gocv.Net
	scale  float64
	resize image.Point
	mean   gocv.Scalar
	swapRB bool
	proc   Processor
}

func NewDetector(conf *Config) Detector {

	log.Printf("Found OpenCV: %s\n", gocv.Version())

	if _, err := os.Stat(conf.ModelPath); os.IsNotExist(err) {
		log.Panicf("Model not found [%s]: %s\n", conf.ModelPath, err.Error())
	}

	if _, err := os.Stat(conf.ConfigPath); os.IsNotExist(err) {
		log.Panicf("Net config not found [%s]: %s\n", conf.ConfigPath, err.Error())
	}

	// create DNN
	net := gocv.ReadNet(conf.ModelPath, conf.ConfigPath)
	if net.Empty() {
		log.Panicf("Error reading network: model:%s,  config:%s\n", conf.ModelPath, conf.ConfigPath)
	}

	if err := net.SetPreferableBackend(conf.Backend); err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}

	if err := net.SetPreferableTarget(conf.Target); err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}

	var proc Processor
	if conf.Type == NetTypeSSD || conf.Type == NetTypeEf {
		proc = NewProcessorSSD()
	} else if conf.Type == NetTypeYOLO4 || conf.Type == NetTypeYOLO3 {
		proc = NewProcessorYOLO()
	} else {
		log.Panicf("Unknown processor\n")
	}

	d := &detector{}
	d.net = &net
	d.proc = proc
	d.resize = conf.Resize
	d.scale = conf.Scale
	d.mean = conf.Mean
	d.swapRB = conf.SwapRB

	return d
}

func (d *detector) Detect(img gocv.Mat, threshold float32, all bool) (boxes []Box) {

	if img.Empty() {
		log.Printf("libdetector: empty frame skipped\n")
		return boxes
	}

	// create resized blob
	blob := gocv.BlobFromImage(img, d.scale, d.resize, d.mean, d.swapRB, false)

	boxes = d._detectBlob(&blob, threshold, all)

	if err := blob.Close(); err != nil {
		log.Printf("[WARN] libdetector: %s", err.Error())
	}

	return boxes
}

func (d *detector) DetectBlob(blob gocv.Mat, threshold float32, all bool) (boxes []Box) {

	boxes = d._detectBlob(&blob, threshold, all)

	if err := blob.Close(); err != nil {
		log.Printf("[WARN] libdetector: %s", err.Error())
	}

	return boxes
}

func (d *detector) _detectBlob(blob *gocv.Mat, threshold float32, all bool) (boxes []Box) {

	// set net input
	d.net.SetInput(*blob, "")

	// start forwarding
	prob := d.net.Forward("")

	boxes = d.proc.Process(prob, threshold, all)

	if err := prob.Close(); err != nil {
		log.Printf("[WARN] libdetector: %s\n", err.Error())
	}

	return boxes
}

func (d *detector) Resize() image.Point {
	return d.resize
}

func (d *detector) Scale() float64 {
	return d.scale
}

func (d *detector) Mean() gocv.Scalar {
	return d.mean
}

func (d *detector) SwapRB() bool {
	return d.swapRB
}
