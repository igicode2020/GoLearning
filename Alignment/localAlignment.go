package main

type Alignment [2]string

// LocalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score local alignment of the strings corresponding to these penalties.
func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	var a Alignment

	table := LocalScoreTable(str1, str2, match, mismatch, gap)
	row, col := MaxMatrix(table)
	meu := mismatch
	sig := gap
	endRow := row
	endCol := col

	for row != 0 && col != 0 {
		diag := table[row-1][col-1]
		if str1[row-1] == str2[col-1] {
			diag += match
		} else {
			diag -= meu
		}
		dir := Max4(table[row-1][col]-sig, table[row][col-1]-sig, diag)
		if dir == 1 {
			a[0] = string(str1[row-1]) + a[0]
			a[1] = "-" + a[1]
			row--
		} else if dir == 0 {
			a[0] = string(str1[row-1]) + a[0]
			a[1] = string(str2[col-1]) + a[1]
			row--
			col--
		} else if dir == -1 {
			a[0] = "-" + a[0]
			a[1] = string(str2[col-1]) + a[1]
			col--
        } else {
			break
		}
	}
	return a, row, endRow, col, endCol
}

func Max4(up, left, diag float64) float64 {
	if up > left && up > diag && up > 0{
		return 1
	} else if diag > left && diag > 0 {
		return 0
    } else if left > 0 {
		return -1
	} else {
		return -2
	}
}

func MaxMatrix(list [][]float64)(int, int) {
    m := list[0][0]
    var row int
    var col int
    for i, val := range list{
        for j, vali := range val {
            if vali > m{
                row = i
                col = j
                m= vali
            }
        }
    }
    retur