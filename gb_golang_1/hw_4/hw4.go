package hw_4

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers = []int{4, 1, 7, 0, 44, 2, 56, 0, 20}
	var result = InsertionSort(numbers)
	for i, r := range result {
		fmt.Println("index = ", strconv.Itoa(i), strconv.Itoa(r))
	}
}

func InsertionSort(ar []int) []int {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
	return ar
}
