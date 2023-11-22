package main

import (
	"fmt"
)

// Error type for negative float64 numbers
type ErrNegativeSqrt float64

// Function for type ErrNegativeSqrt to return appropriate error message
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Function to calculate Square root
func Sqrt(x float64) (float64, error) {

	// If input number is negative return error
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.421
	z -= (z*z - x) / (2 * z)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
