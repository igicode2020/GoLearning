package main

//EditDistanceMatrix takes as input a slice of strings patterns.
//It returns a matrix whose (i,j)th entry is the edit distance between
//the i-th and j-th strings in patterns.
func EditDistanceMatrix(patterns []string) [][]int {
	table := make([][]int, len(patterns))
	for i := range table {
        table[i] = make([]int, len(patterns))
    }
	for j := range table {
		for v := j+1; v < len(patterns); v++ {
			table[j][v] = EditDistance(patterns[j], patterns[v])
			table[v][j] = table[j][v]
		}
	}
	return table
}
