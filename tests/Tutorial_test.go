package tests

import (
	"testing"
)


func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiply(99999999999999999, 99999999999999999)
	}
}


func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		power(99999999999999999, 99999999999999999)
	}
}
