package goreco

import (
	"bytes"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func MarkObjects(img *gocv.Mat, boxes Boxes, putText bool) {
	scaleX := float32(img.Cols())
	scaleY := float32(img.Rows())

	for i := 0; i < len(boxes); i += 1 {
		box := boxes[i]

		scaledRect := image.Rect(
			int(box.rect.min[0]*scaleX), int(box.rect.min[1]*scaleY),
			int(box.rect.max[0]*scaleX), int(box.rect.max[1]*scaleY),
		)
		gocv.Rectangle(img, scaledRect, color.RGBA{0, 255, 0, 0}, 2)

		if putText {
			gocv.PutText(img, fmt.Sprintf("%d(%f)", box.class, box.conf), scaledRect.Min, 1, 2, color.RGBA{100, 200, 0, 0}, 3)
		}
	}
}

type Jpeg []byte

var JpegHead = []byte{0xff, 0xd8}
var JpegTail = []byte{0xFF, 0xD9}

func CheckJpegHeader(data []byte) bool {
	return bytes.Equal(data[:2], JpegHead)
	//return bytes.Equal(data[:2], JpegHead) && bytes.Equal(data[len(data)-2:], JpegTail)
}

func FindMaxFloat32(floats []float32) (max float32, idx int) {
	for i, v := range floats {
		if v > max {
			max = v
			idx = i
		}
	}

	return max, idx
}
