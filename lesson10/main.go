package main

import (
	"fmt"
	"lesson10/fibonacci"
	"lesson10/primenum"
	"lesson10/sort"
)

func main() {
	var stopNumber int
	fmt.Print("Enter your number: ")
	_, err := fmt.Scan(&stopNumber)
	if err != nil {
		fmt.Println("Incorrect input")
	} else {
		fmt.Println("Prime Numbers are", primenum.GetPrimes(stopNumber))
		fmt.Println("Prime Numbers (Naive) are", primenum.GetPrimesNaive(stopNumber))
		fmt.Printf("Fibonacci number %d has value %d\n", stopNumber, fibonacci.FibWithCache(stopNumber))

	}
	fmt.Println("Some sorted slice is", sort.InsertionSort([]int{5, 2, 10, 6, 100, 54, 919, 351, 11, 523, 63, 122, 473, 890, 1000}))
}
