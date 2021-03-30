package libdetector

import (
	"gocv.io/x/gocv"
	"image"
)

type NetType int

const (
	NetTypeSSD NetType = iota
	NetTypeYOLO3
	NetTypeYOLO4
	NetTypeYOLO5
	NetTypeEf
)

type Config struct {
	Type       NetType
	ModelPath  string
	ConfigPath string
	Resize     image.Point
	Scale      float64
	Mean       gocv.Scalar
	SwapRB     bool
	Backend    gocv.NetBackendType
	Target     gocv.NetTargetType
}

var DefaultSSDConfig = &Config{
	NetTypeSSD,
	"",
	"",
	image.Pt(300, 300),
	1.0 / 127.5,
	gocv.NewScalar(127.5, 127.5, 127.5, 0),
	true,
	gocv.NetBackendDefault,
	gocv.NetTargetCPU,
}

var DefaultEfConfig = &Config{
	NetTypeEf,
	"",
	"",
	image.Pt(512, 512),
	1.0 / 127.5,
	gocv.NewScalar(127.5, 127.5, 127.5, 0),
	true,
	gocv.NetBackendDefault,
	gocv.NetTargetCPU,
}

var DefaultYOLO3Config = &Config{
	NetTypeYOLO3,
	"",
	"",
	image.Pt(320, 320),
	1.0 / 255.0,
	gocv.NewScalar(255.0, 255.0, 255.0, 0),
	true,
	gocv.NetBackendDefault,
	gocv.NetTargetCPU,
}

var DefaultYOLO4Config = &Config{
	NetTypeYOLO3,
	"",
	"",
	image.Pt(608, 608),
	1.0 / 255.0,
	gocv.NewScalar(255.0, 255.0, 255.0, 0),
	true,
	gocv.NetBackendDefault,
	gocv.NetTargetCPU,
}
