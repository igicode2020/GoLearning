package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your MaxIntegerArray() function here, along with any subroutines that you need.
func MaxIntegerArray(list []int) int {
	max := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}
	return max
}
