package hw_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonachi(t *testing.T) {
	var result = fibbonachi(4)
	asserts := assert.New(t)
	asserts.Equal(3, result, "Not match")
}

func BenchmarkFibonachiRecursion(b *testing.B) {
	println("BenchmarkFibonachiRecursion")

	for i := 0; i < b.N; i++ {
		fibbonachi(30)
	}
}

func BenchmarkFibonachiMaps(b *testing.B) {
	println("BenchmarkFibonachiMaps")

	for i := 0; i < b.N; i++ {
		fibbonachiWithMap(30)
	}
}
