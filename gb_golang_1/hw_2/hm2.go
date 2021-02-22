package main

//package hw_2

import (
	"fmt"
	"strconv"
)

func main() {
	calculateSquareOfRectangle()
	calculateDiameterAndLengthOfCircleSquare()
	calculate3TypeNumber()
}

func calculateSquareOfRectangle() {
	var a, b int
	fmt.Print("Введите первое число:")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число:")
	fmt.Scanln(&b)

	var square = a * b

	fmt.Print("Площадь прямоугольника = " + strconv.Itoa(square) + "\n")
}

func calculateDiameterAndLengthOfCircleSquare() {
	var a float64
	fmt.Print("Введите площадь круга:")
	fmt.Scanln(&a)

	var temp = a * 4 * 3.14
	var lengthOfCircle int64

	for n := int32(3); n < int32(temp); n += 2 {
		lengthOfCircle = int64(n)
		s := lengthOfCircle * lengthOfCircle
		for i := int64(3); i*i < s; i += 2 {
			if lengthOfCircle%i == 0 {
				break
			}
		}
	}

	fmt.Println("Длина окружности = " + strconv.Itoa(int(lengthOfCircle)))

}

func calculate3TypeNumber() {
	var a int
	fmt.Print("Введите число:")
	fmt.Scanln(&a)

	var s string = strconv.Itoa(a)

	fmt.Print("Результат: \nСотых " + string(s[0]) + "\nДесятичных " + string(s[1]) + "\nЕдиниц " + string(2))
}
