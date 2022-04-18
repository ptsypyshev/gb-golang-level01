package primenum

import (
	"fmt"
	"reflect"
	"testing"
)

var resultMap = map[int][]int{
	0:  nil,
	1:  nil,
	2:  {2},
	3:  {2, 3},
	4:  {2, 3},
	5:  {2, 3, 5},
	6:  {2, 3, 5},
	7:  {2, 3, 5, 7},
	8:  {2, 3, 5, 7},
	9:  {2, 3, 5, 7},
	10: {2, 3, 5, 7},
}

func TestGetPrimes(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Run(fmt.Sprintf("Optimized func with case with n = %d", i), func(t *testing.T) {
			if !reflect.DeepEqual(GetPrimes(i), resultMap[i]) {
				t.Errorf("Incorrect result for n = %d", i)
			}
		})
	}
}

func TestGetPrimesNaive(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Run(fmt.Sprintf("Naive func with case with n = %d", i), func(t *testing.T) {
			if !reflect.DeepEqual(GetPrimesNaive(i), resultMap[i]) {
				t.Errorf("Incorrect result for n = %d", i)
			}
		})
	}
}

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(10)
	}
}

func BenchmarkGetPrimesNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimesNaive(10)
	}
}

func ExampleGetPrimesNaive() {
	f := GetPrimesNaive(10)
	fmt.Println(f)

	//Output: [2, 3, 5, 7]
}

func ExampleGetPrimes() {
	f := GetPrimes(10)
	fmt.Println(f)

	//Output: [2, 3, 5, 7]
}
