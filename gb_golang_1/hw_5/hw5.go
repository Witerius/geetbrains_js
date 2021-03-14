package hw_5

import "fmt"

var fibbonachiMap = make(map[int]int)

func main() {
	fmt.Println("Простой")
	fmt.Println(fibbonachi(0))
	fmt.Println(fibbonachi(1))
	fmt.Println(fibbonachi(2))
	fmt.Println(fibbonachi(3))
	fmt.Println(fibbonachi(4))
	fmt.Println(fibbonachi(5))

	fmt.Println("Оптимизированный")

	for i := 0; i < 50; i++ {
		fmt.Println(fibbonachiWithMap(i))
	}

}

func fibbonachi(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}

func fibbonachiWithMap(n int) int {

	switch {
	case n == 0:
		fibbonachiMap[0] = 0
	case n == 1:
		fibbonachiMap[1] = 1
	default:
		fibbonachiMap[n] = fibbonachiMap[n-1] + fibbonachiMap[n-2]
	}

	return fibbonachiMap[n]
}
