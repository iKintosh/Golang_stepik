package main

import "math"

var k float64
var p float64
var v float64

func M() float64 {
	return p * v
}

func W() float64 {
	return math.Sqrt(k/M())
}

func T() float64 {
	return 6.0/W()
}