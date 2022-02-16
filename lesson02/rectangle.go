package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("Enter one rectangle side: ")
	_, err1 := fmt.Scan(&a)
	fmt.Print("Enter another rectangle side: ")
	_, err2 := fmt.Scan(&b)
	if err1 != nil || err2 != nil {
		fmt.Println("Incorrect input")
	} else {
		fmt.Println("The area of a rectangle is", a*b)
	}
}
