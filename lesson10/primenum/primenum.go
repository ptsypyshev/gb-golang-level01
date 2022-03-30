package primenum

//GetPrimesNaive is a function that returns a slice of prime numbers up to input value. Naive version of func.
func GetPrimesNaive(n int) (res []int) {
MainLoop:
	for i := 2; i <= n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue MainLoop
			}
		}
		res = append(res, i)
	}
	return
}

//GetPrimes is a function that returns a slice of prime numbers up to input value. Optimized version of func.
func GetPrimes(n int) (res []int) {
	//Init Slice with true values for Sieve of Eratosthenes
	//https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%88%D0%B5%D1%82%D0%BE_%D0%AD%D1%80%D0%B0%D1%82%D0%BE%D1%81%D1%84%D0%B5%D0%BD%D0%B0
	if n < 2 {
		return res
	}
	initSlice := make([]bool, n+1)

	// Start Searching Prime Numbers
	step := 2
	for step <= n {
		if !initSlice[step] {
			for j := 2 * step; j <= n; j += step {
				initSlice[j] = true // cross a composite number off the slice
			}
		}
		step += 1
	}

	for idx, val := range initSlice[2:] {
		if !val {
			res = append(res, idx+2)
		}
	}
	return
}
