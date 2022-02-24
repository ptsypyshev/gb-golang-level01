package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter your integer number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Incorrect input")
	} else {
		if number >= 1000 || number <= -1000 {
			fmt.Println("Your number is out of range (not more than 3 digits)")
		} else {
			hundreds := number / 100
			dozens := number % 100 / 10
			units := number % 10
			fmt.Println("Your number consist of", hundreds, "hundreds,", dozens, "dozens,", units, "units.")
		}
	}
}
