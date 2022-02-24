package main

import "fmt"

func main() {
	var stopNumber int
	fmt.Print("Enter your number: ")
	_, err := fmt.Scan(&stopNumber)
	if err != nil {
		fmt.Println("Incorrect input")
	} else {
		//Init Slice with true values for Sieve of Eratosthenes
		//https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%88%D0%B5%D1%82%D0%BE_%D0%AD%D1%80%D0%B0%D1%82%D0%BE%D1%81%D1%84%D0%B5%D0%BD%D0%B0
		initSlice := make([]bool, stopNumber+1)

		// Start Searching Prime Numbers
		step := 2
		for step <= stopNumber {
			if !initSlice[step] {
				for j := 2 * step; j <= stopNumber; j += step {
					initSlice[j] = true // cross a composite number off the slice
				}
			}
			step += 1
		}

		fmt.Print("Prime Numbers are ")
		for idx, val := range initSlice[2:] {
			if !val {
				fmt.Printf("%v ", idx+2)
			}
		}
	}
}
