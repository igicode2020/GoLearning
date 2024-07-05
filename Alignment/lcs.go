package main

import "strings"

// LongestCommonSubsequence takes two strings as input.
// It returns a longest common subsequence of the two strings.
func LongestCommonSubsequence(str1, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	strings.Split(str1, "")
	strings.Split(str2, "")
	u := len(str1)
	p := len(str2)
	//make table
	table := make([][]int, len(str1)+1)
	for i := range table {
		table[i] = make([]int, len(str2)+1)
	}
	backtrack := make([][]int, len(str1)+1)
	for i := range backtrack {
		backtrack[i] = make([]int, len(str2)+1)
	}

	for n := 1; n < len(table); n++ {
		for k := 1; k < len(table[n]); k++ {
			diag := table[n-1][k-1]
			if str1[n-1] == str2[k-1] {
				diag++
			}
			table[n][k], backtrack[n][k] = Max3(table[n-1][k], table[n][k-1], diag)
		}
	}
	//make reversed list
	revList := make([]byte, 0, table[len(str1)][len(str2)])

	// if -1 left; if 0 diag; if 1 up
	for u > 0 && p > 0 {
		if backtrack[u][p] == 0 {
			if str1[u-1] == str2[p-1] {
				revList = append(revList, str1[u-1])
			}
			u--
			p--
		} else if backtrack[u][p] == 1 {
			u--
		} else if backtrack[u][p] == -1 {
			p--
		}
	}
	list := Reverse(string(revList))

	return list
}
func Max3(up, left, diag int) (int, int) {
	if up > left && up > diag {
		return up, 1
	} else if diag > left {
		return diag, 0
	}
	return left, -1
}

func Reverse(s string) string {
	and := make([]byte, len(s))
	for i := range s {
		and[i] = s[len(s)-1-i]
	}
	return string(and)
}
