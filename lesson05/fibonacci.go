package main

import "fmt"

func fibonacci(n int, fibMap map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if fibMap == nil {
		fibMap = make(map[int]int, n)
		fibMap[0] = 0
		fibMap[1] = 1
	}

	a, existA := fibMap[n-2]
	b, existB := fibMap[n-1]

	if !existB {
		b = fibonacci(n-1, fibMap)
		fibMap[n-1] = b
	}

	if !existA {
		a = fibonacci(n-2, fibMap)
		fibMap[n-2] = a
	}

	return a + b
}

func main() {
	var n int
	fmt.Print("Enter Fibonacci number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Enter a digit!")
	} else {
		fmt.Printf("Value of your Fibonacci number is %d", fibonacci(n, nil))
	}

}
