package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your DividesAll() function here, along with any subroutines that you need.
func DividesAll(a []int, d int) bool {
    if d == 0 {
        return false
    }
    sumTest := 0
    for i := 0; i < len(a); i++ {
        if a[i] % d == 0 {
            sumTest += 1
        }
    }
    if sumTest == len(a) {
        return true
    }
    
    return false

}