package goplayground

import (
	"fmt"
	"log"
	"math"
)

// Sqrt calculates the square root of the given float.
// It will return 0 and error if given value is <0.
func Sqrt(f float64) (float64, error) {
	if f < 0.0 {
		// Error message should not end in a punctuation
		return 0, fmt.Errorf("Input Val is < 0")
	}
	return math.Sqrt(f), nil
}

// useSqrt is an example of using the Sqrt function above.
func useSqrt() {
	//root, err := sqrt(-3.0)
	root, err := Sqrt(3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%.5f\n", root)
	}
}
