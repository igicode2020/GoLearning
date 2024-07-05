package main

import "strings"

// LCSLength takes two strings as input. It returns the length of a longest common
// subsequence of the two strings.
func LCSLength(str1, str2 string) int {
	strings.Split(str1, "")
	strings.Split(str2, "")

	table := make([][]int, len(str1))

	for i := range table {
		table[i] = make([]int, len(str2))
	}

	for n := 1; n < len(table); n++ {
		for k := 1; k < len(table[n]); n++ {
			table[n][k] = Max3(table[n-1][k], table[n][k-1], table[n - 1][k - 1], str1[n], str2[k])
		}
	}
	return table[len(str1) - 1][len(str2) - 1]
}

func Max3(a, b, c int, str1, str2 string) int {
    if a >= b || c >= b {
        if a > c {
            return a
        }
		if str1 == str2 {
        return c + 1
		} else {
			return c
		}
    }
    return b
}