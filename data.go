package libdetector

import "gocv.io/x/gocv"

type RecDataType int

const (
	TypeRecDataBytes = iota
	TypeRecDataJpeg
)

type RecData interface {
	GetType() RecDataType
	GetData() []byte
}

type RecDataBytes interface {
	RecData
	Cols() int
	Raws() int
	Mt() gocv.MatType
}

type RecDataJpeg interface {
	RecData
}

type RecDataJ struct {
}

type recData struct {
	recDataType RecDataType
	data        []byte
}

type recDataB struct {
	rows, cols int
	mt         gocv.MatType
}

func NewRecBytesJpeg(data []byte) RecData {
	return struct {
		*recData
	}{
		&recData{TypeRecDataJpeg, data},
	}
}

func NewRecBytesData(data []byte, cols, rows int, mt gocv.MatType) RecData {
	return struct {
		*recData
		*recDataB
	}{
		&recData{TypeRecDataBytes, data},
		&recDataB{rows, cols, mt},
	}
}

func (r *recData) GetType() RecDataType {
	return r.recDataType
}

func (r *recData) GetData() []byte {
	return r.data
}

func (r *recDataB) Cols() int {
	return r.Cols()
}

func (r *recDataB) Rows() int {
	return r.rows
}

func (r *recDataB) Mt() gocv.MatType {
	return r.mt
}
