package calculator

type Calculator interface {
	 Add() int
}

type simpleClr struct {
	x, y int
}

func NewCalculator(x, y int) Calculator {
	return &simpleClr{x, y}
}

func (s *simpleClr) Add() int {
	return s.x + s.y
}