package main

type Alignment [2]string

// GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var a Alignment
	//setup global score
	table := GlobalScoreTable(str1, str2, match, mismatch, gap)
	meu := mismatch
	sig := gap

	row := len(str1)
	col := len(str2)

	for row != 0 && col != 0 {
		diag := table[row-1][col-1]
		if str1[row-1] == str2[col-1] {
			diag += match
		} else {
			diag -= meu
		}
		dir := Max3(table[row-1][col]-sig, table[row][col-1]-sig, diag)
		if dir == 1 {
			a[0] = string(str1[row-1]) + a[0]
			a[1] = "-" + a[1]
			col--
		} else if dir == 0 {
			a[0] = string(str1[row-1]) + a[0]
			a[1] = string(str2[col-1]) + a[1]
			row--
			col--
		} else {
			a[0] = "-" + a[0]
			a[1] = string(str2[col-1]) + a[1]
			row--
		}
	}

	return a
}

func Max3(up, left, diag float64) float64 {
	if up > left && up > diag {
		return 1
	} else if diag > left {
		return 0
	}
	return -1
}
