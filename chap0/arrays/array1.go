package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your FibonacciArray() function here, along with any subroutines that you need.
func FibonacciArray(n int) []int {
    if n < 0 {
        return nil
	}

	// Create a slice to store Fibonacci numbers
	fib := make([]int, n+1)

	// Base cases
	if n >= 0 {
		fib[0] = 1
	}
	if n >= 1 {
		fib[1] = 1
	}

	// Calculate Fibonacci numbers and store in the slice
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}