package main

// GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	//declare weight values
	meu := mismatch
	sig := gap

	table := make([][]float64, len(str1)+1)
	for i := range table {
		table[i] = make([]float64, len(str2)+1)
	}

	for n := 1; n < len(table); n++ {
		table[n][0] = table[n-1][0] + sig
	}

	for k := 1; k < len(table[0]); k++ {
		table[0][k] = table[0][k-1] + sig
	}

	for n := 1; n < len(table); n++ {
		for k := 1; k < len(table[n]); k++ {
			diag := table[n-1][k-1]
			if str1[n-1] == str2[k-1] {
				diag += match
			} else {
				diag -= meu
			}
			table[n][k] = Max3(table[n-1][k]-sig, table[n][k-1]-sig, diag)
		}
	}
	return table

}
func Max3(up, left, diag float64) float64 {
	if up > left && up > diag {
		return up
	} else if diag > left {
		return diag
	}
	return left
}
