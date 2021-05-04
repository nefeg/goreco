# goreco

OpenCV-DNN wrapper written in GO for using in Unspot.

# Dependencies
- [OpenCV](https://opencv.org/)

# Supported platforms

 - Linux
 - OS X

# Getting started
## How to use
```shell
go get "github.com/umbrella-evgeny-nefedkin/goreco"
```

```go
package main

import (
	"github.com/umbrella-evgeny-nefedkin/goreco"
        "gocv.io/x/gocv"
)

func processFrame(img gocv.Mat) Boxes{
    // Create pre-config
	Config := goreco.DefaultSSDConfig

    // Set model path
	Config.ModelPath = "model/ssd_mobile/frozen_inference_graph.pb"
	Config.ConfigPath = "model/ssd_mobile/ssd.pbtxt"

    // Create detector
    Detector := goreco.NewDetector(Config)

    // Call detection
    result := Detector.Detect(img, 0.6)

    // Mark detected regions on original frame
    goreco.MarkObjects(&img, boxes, true)

    return result
}
```
