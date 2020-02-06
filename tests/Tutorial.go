package tests

import (
	"math"
	"time"
)

type Tutorial struct {

}

func add(i1, i2 int) int {
	return i1 + i2
}

func multiply(i1, i2 int) int {
	return i1 * i2
}


func power(i1, i2 float64) float64 {
	time.Sleep(1 * time.Millisecond)
	return math.Pow(i1, i2)
}
