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
		initSlice := make([]bool, stopNumber)

		// Start Searching Prime Numbers
		var result []int
		step := 2
		for step < stopNumber {
			result = append(result, step)
			for j := 2 * step; j < stopNumber; j += step {
				initSlice[j] = true // cross a composite number off the slice
			}
			nextStep := step
			for idx, val := range initSlice[step+1:] {
				// Check the next true element in slice
				if !val {
					nextStep += idx + 1
					break
				}
			}

			// Break the loop if no next step found
			if nextStep == step {
				break
			} else {
				step = nextStep
			}
		}
		fmt.Println("Prime Numbers are", result)
	}
}
