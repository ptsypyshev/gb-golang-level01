package sort

//InsertionSort is a simple sorting algorithm. It gets unsorted slice and returns a sorted copy of the slice.
func InsertionSort(origSlice []int) []int {
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

//BubbleSort is also a sorting algorithm. It gets unsorted slice and returns a sorted copy of the slice.
func BubbleSort(origSlice []int) []int {
	lenSlice := len(origSlice) - 1
	res := make([]int, lenSlice+1)
	copy(res, origSlice)

	for i := 0; i < lenSlice; i++ {
		for j := 0; j < lenSlice-i; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	return res
}
