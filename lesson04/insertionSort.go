package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(origSlice []int) []int {
	lenSlice := len(origSlice)
	res := make([]int, lenSlice)
	copy(res, origSlice)

	for i := 1; i < lenSlice; i++ {
		for j := i; j > 0 && res[j] < res[j-1]; j-- {
			res[j], res[j-1] = res[j-1], res[j]
		}
	}
	return res
}

func main() {
	fmt.Print("Enter Unsorted Sequence: ")
	input := bufio.NewReader(os.Stdin)
	line, err := input.ReadString('\n')

	if err != nil {
		fmt.Println("Something went wrong!")
		return
	}

	stringSlice := strings.Fields(line)
	unsortedSlice := make([]int, 0, len(stringSlice))
	for _, elem := range stringSlice {
		intElem, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Incorrect input (not only digits)!")
			os.Exit(1)
		}
		unsortedSlice = append(unsortedSlice, intElem)
	}
	sortedSlice := insertionSort(unsortedSlice)
	fmt.Print("Sorted Sequence is: ")
	for _, elem := range sortedSlice {
		fmt.Printf("%d ", elem)
	}
}
