package main

import "strings"

// EditDistance takes two strings as input. It returns the Levenshtein distance
// between the two strings; that is, the minimum number of substitutions, insertions, and deletions
// needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	strings.Split(str1, "")
	strings.Split(str2, "")

	table := make([][]int, len(str1)+1)

	for i := range table {
		table[i] = make([]int, len(str2)+1)
	}

	for n := 1; n < len(table); n++ {
		table[n][0] = table[n-1][0] + 1
	}

	for k := 1; k < len(table[0]); k++ {
		table[0][k] = table[0][k-1] + 1
	}

	for n := 1; n < len(table); n++ {
		for k := 1; k < len(table[n]); k++ {
			up := table[n-1][k] + 1
			left := table[n][k-1] + 1
			diag := table[n-1][k-1] + 1
			if str1[n-1] == str2[k-1] {
				diag--
			}
			table[n][k] = Min3(up, left, diag)
		}
	}
	return table[len(str1)][len(str2)]
}

func Min3(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if c < b {
		return c
	}
	return b
}
