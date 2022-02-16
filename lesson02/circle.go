package main

import (
	"fmt"
	"math"
)

func main() {
	var circleArea float64
	fmt.Print("Enter area of circle: ")
	_, err := fmt.Scan(&circleArea)
	if err != nil {
		fmt.Println("Incorrect input")
	} else {
		radius := math.Sqrt(circleArea / math.Pi)
		diameter := 2 * radius
		circumference := math.Pi * diameter
		fmt.Println("Diameter of this circle is", diameter, "circumference is", circumference)
	}
}
