package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your Combination() function here, along with any subroutines that you need.
func Combination(n, k int) int {
    factN := 1
    subtractKN := n - k
    factKN := 1
    factK := 1
    for n > 0 {
        factN = factN * n
        n -= 1
    }
    for subtractKN > 0 {
        factKN = factKN * subtractKN
        subtractKN -= 1
    }
    for k > 0 {
        factK = factK * k
        k -= 1
    }
    return factN / (factKN * factK)
}
