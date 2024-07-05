package main

// ALL PENALTIES POSITIVE
var a Alignment

// ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
// It returns the maximum score of an overlap alignment in which str1 is overlapped with str2.
// Assume we are overlapping a suffix of str1 with a prefix of str2.
func ScoreOverlapAlignment(str1, str2 string, match, mismatch, gap float64) float64 {
	table := LocalScoreTable(str1, str2, match, mismatch, gap)
	max := 0.0
	for i := range table {
		if table[i][len(table[i])-1] > max {
			max = table[i][len(table[i])-1]
		}
	}
	return max
}

func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {

	meu := mismatch
	sig := gap

	table := make([][]float64, len(str1)+1)
	for i := range table {
		table[i] = make([]float64, len(str2)+1)
	}
	for k := range table {
		table[0][k] = -(gap * k)
	}

	for n := 1; n < len(table); n++ {
		for k := 1; k < len(table[n]); k++ {
			diag := table[n-1][k-1]
			if str1[n-1] == str2[k-1] {
				diag += match
			} else {
				diag -= meu
			}
			table[n][k] = Max4(table[n-1][k]-sig, table[n][k-1]-sig, diag)
		}
	}

	return table
}

func Max4(up, left, diag float64) float64 {
	if up > left && up > diag && up > 0 {
		return up
	} else if diag > left && diag > 0 {
		return diag
	} else if left > 0 {
		return left
	} else {
		return 0
	}
}
