package goreco

import "log"

type Stack interface {
	Push(int)
	Size() uint
	SizeCur() uint
	Mean() uint
	Sum() int
	Avg() float64
	AvgNoZero() int
	Max() int
}

type stack struct {
	size uint
	rr   []int
}

//NewResultDamper
/*
 * size - stack size
 */
func NewStack(size uint) Stack {
	if size == 0 {
		log.Panicln("size must be greater then 0")
	}
	return &stack{size: size}
}

//Push - push value in the stack
func (s *stack) Push(r int) {
	if uint(len(s.rr)) >= s.size {
		s.rr = s.rr[1:]
	}

	s.rr = append(s.rr, r)
}

//Size - return declared size of the stack
func (s *stack) Size() uint {
	return s.size
}

//SizeCur - return CURRENT size of the stack
func (s *stack) SizeCur() uint {
	return uint(len(s.rr))
}

//Mean - amount of non-zero values
func (s *stack) Mean() uint {
	var mean uint = 0
	for i := 0; i < len(s.rr); i++ {
		if s.rr[i] > 0 {
			mean += 1
		}
	}

	return mean
}

//Sum - sum of values in the stack
func (s *stack) Sum() int {
	sum := 0
	for i := 0; i < len(s.rr); i++ {
		sum += s.rr[i]
	}

	return sum
}

//Avg - average mean of the stack
func (s *stack) Avg() float64 {
	panic("implement me")
}

//AvgNoZero - average mean of non-zero values in the stack
func (s *stack) AvgNoZero() int {
	panic("implement me")
}

//Max - max value of the stack
func (s *stack) Max() int {

	var max int
	for i, e := range s.rr {
		if i == 0 || e > max {
			max = e
		}
	}

	return max
}
