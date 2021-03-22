package hw_4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var table = []struct {
	in  []int
	out []int
}{
	{[]int{4, 1, 7, 0, 44, 2, 56, 20}, []int{0, 1, 2, 4, 7, 20, 44, 56}},
	{[]int{5, 2, 3}, []int{2, 3, 5}},
	{[]int{6, 0, 1}, []int{0, 1, 6}},
}

func TestTableDriven(t *testing.T) {
	var position int = 0
	for _, tt := range table {

		t.Run(string(rune(position)), func(t *testing.T) {
			position += 1
			var result = InsertionSort(tt.in)

			asserts := assert.New(t)
			asserts.Equal(tt.out, result, "Not match")
		})
	}

}

func TestInsertionSort(t *testing.T) {
	var expected = []int{0, 1, 2, 4, 7, 20, 44, 56}
	var params = []int{4, 1, 7, 0, 44, 2, 56, 20}
	var result = InsertionSort(params)

	asserts := assert.New(t)
	asserts.Equal(expected, result, "Not match")
}

func ExampleInsertionSort() {
	var params = []int{4, 1, 7, 0, 44, 2, 56, 20}
	result := InsertionSort(params)
	fmt.Print(result)
	// Output: [0 1 2 4 7 20 44 56]
}
