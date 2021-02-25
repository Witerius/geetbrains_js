package main

import (
	"fmt"
)

//package hw_3

func main() {
	var a, b float32
	var operator string

	fmt.Print("Введите первое число:")
	fmt.Scanln(&a)

	fmt.Print("Введите оператор:")
	fmt.Scanln(&operator)

	fmt.Print("Введите второе число:")
	fmt.Scanln(&b)

	switch operator {
	case "+":
		casePlus(a, b)
	case "-":
		caseMinus(a, b)
	case "/":
		caseDivision(a, b)
	case "*":
		caseMultiplicating(a, b)
	}
}

func caseMinus(a float32, b float32) {
	var result = a - b
	textResult := fmt.Sprintf("Результат = %f", result)
	fmt.Println(textResult)
}

func casePlus(a float32, b float32) {
	var result = a + b
	textResult := fmt.Sprintf("Результат = %f", result)
	fmt.Println(textResult)
}

func caseDivision(a float32, b float32) {
	if b == 0 {
		fmt.Println("Нельзя делить на ноль")
		return
	}
	var result = a / b
	textResult := fmt.Sprintf("Результат = %f", result)
	fmt.Println(textResult)
}

func caseMultiplicating(a float32, b float32) {
	var result = a * b
	textResult := fmt.Sprintf("Результат = %f", result)
	fmt.Println(textResult)
}
