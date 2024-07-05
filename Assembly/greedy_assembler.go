package main

// GreedyAssembler takes a collection of strings and returns a genome whose
// k-mer composition is these strings. It makes the following assumptions.
// 1. "Perfect coverage" -- every k-mer is detected
// 2. No errors in reads
// 3. Every read has equal length (k)
// 4. DNA is single-stranded
func GreedyAssembler(reads_ []string) string {
	reads := reads_
	genome := reads[0]
	k := len(genome)
	reads = reads[1:]
	ind := CheckSuffixes(reads, genome[:k-1])
	for ind != -1 {
		genome = string(reads[ind][0]) + genome
		reads = append(reads[:ind], reads[ind+1:]...)
		ind = CheckSuffixes(reads, genome[:k-1])
	}
	ind = CheckPrefixes(reads, genome[len(genome)-k+1:len(genome)])
	for ind != -1 {
		genome = genome + string(reads[ind][k-1])
		reads = append(reads[:ind], reads[ind+1:]...)
		ind = CheckPrefixes(reads, genome[len(genome)-k+1:len(genome)])
	}
	return genome
}

func CheckSuffixes(reads []string, relstr string) int {
	for i, val := range reads {
		if val[1:] == relstr {
			return i
		}
	}
	return -1
}

func CheckPrefixes(reads []string, relstr string) int {
	for i, val := range reads {
		if val[:len(val)-1] == relstr {
			return i
		}
	}
	return -1
}
