package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your CountSharedKmers() function here, along with any subroutines that you need.
func CountSharedKmers(str1, str2 string, k int) int {
	list1 := FrequencyTable(str1, k)
	list2 := FrequencyTable(str2, k)

	return SumOfMinima(list1, list2)
}

// frequency table
func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text) - k
	for i := 0; i <= n; i++ {
		pattern := text[i : i+k]
		freqMap[pattern]++
	}
	return freqMap
}

// minima functions
func SumOfMinima(sample1, sample2 map[string]int) int {
	sum := 0
	for i, count := range sample1 {
		sum += Min2(count, sample2[i])
	}
	return sum
}
func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
