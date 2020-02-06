package tests

import (
	"math"
	"time"
)

type Tutorial struct {

}

func (t Tutorial) add(i1, i2 int) int {
	return i1 + i2
}

func (t Tutorial) multiply(i1, i2 int) int {
	return i1 * i2
}


func (t Tutorial) power(i1, i2 float64) float64 {
	time.Sleep(1 * time.Millisecond)
	return math.Pow(i1, i2)
}
