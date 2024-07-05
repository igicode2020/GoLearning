package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your Permutation() function here, along with any subroutines that you need.
func Permutation(n, k int) int {
    sum0 := 1
    dividend := n-k
    for n > 0 {
        sum0 = sum0 * n
        n -= 1
    }
    sum1 := 1
    for dividend > 0 {
        sum1 = dividend * sum1
        dividend -= 1
    }
    return sum0 / sum1
}