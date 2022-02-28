package main

import (
	"fmt"
	"math"
	"os"
)

func calc(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			os.Exit(1)
		}
		return a / b
	case "^":
		fallthrough
	case "**":
		return math.Pow(a, b)
	case "!":
		intPart, divPart := math.Modf(a)
		if intPart < 0 || divPart > 0 {
			fmt.Println("Факториал вычисляется только для натуральных чисел")
			os.Exit(1)
		}
		resInt := 1
		for i := 1; i <= int(intPart); i++ {
			resInt *= i
		}
		return float64(resInt)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	return 0
}

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	_, err1 := fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	_, err2 := fmt.Scan(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^ или **, !): ")
	_, err3 := fmt.Scan(&op)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Некоректно введены числа или оператор!")
	} else {
		fmt.Printf("Результат выполнения операции: %.2f\n", calc(a, b, op))
	}
}
