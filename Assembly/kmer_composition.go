package main

// KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.
func KmerComposition(genome string, k int) []string {
	array := make([]string, 0, len(genome)-k)
	for i := 0; len(genome) >= i+k; i++ {
		pattern := genome[i : i+k]
		array = append(array, pattern)
	}
	return array
}
