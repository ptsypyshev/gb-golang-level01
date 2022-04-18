package sort

import (
	"fmt"
	"reflect"
	"testing"
)

var origSlice = []int{5, 2, 10, 6, 100, 54, 919, 351, 1000, 11, 523, 63, 122, 473, 890}

var testCases = []struct {
	name   string
	input  []int
	result []int
}{
	{
		name:   "empty slice",
		input:  []int{},
		result: []int{},
	},
	{
		name:   "sorted slice",
		input:  []int{1, 2, 3},
		result: []int{1, 2, 3},
	},
	{
		name:   "negative elements in slice",
		input:  []int{4, -2, 7, 1, -10, 5},
		result: []int{-10, -2, 1, 4, 5, 7},
	},
	{
		name:   "elements up to 1000 in slice",
		input:  []int{5, 2, 10, 6, 100, 54, 919, 351, 1000, 11, 523, 63, 122, 473, 890},
		result: []int{2, 5, 6, 10, 11, 54, 63, 100, 122, 351, 473, 523, 890, 919, 1000},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			if !reflect.DeepEqual(BubbleSort(cse.input), cse.result) {
				t.Errorf("Incorrect result in \"%s\" case", cse.name)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			if !reflect.DeepEqual(InsertionSort(cse.input), cse.result) {
				t.Errorf("Incorrect result in \"%s\" case", cse.name)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(origSlice)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(origSlice)
	}
}

func ExampleBubbleSort() {
	unsortedSlice := []int{10, 2, 6, 1}
	s := BubbleSort(unsortedSlice)
	fmt.Println(s)

	//Output: [1 2 6 10]
}

func ExampleInsertionSort() {
	unsortedSlice := []int{10, 2, 6, 1}
	s := InsertionSort(unsortedSlice)
	fmt.Println(s)

	//Output: [1 2 6 10]
}
