package goplayground

import "fmt"

// FibOne calculates the fibonacci number (slowly), prints the numbers line by line.
// It will also return the final value.
// Note that this function is not optimized whatsoever.
func FibOne(n int) int {
	res := 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 2; i < n; i++ {
			res += FibOne(i - 1)
		}
		fmt.Println(res)
		return res
	}
}

// FibTwo calculates the fibonacci number (slowly).
// It will only return the final value.
// Note that this function is not optimized whatsoever.
func FibTwo(n int) int {
	if n <= 1 {
		return n
	}
	return FibTwo(n-1) + FibTwo(n-2)
}
