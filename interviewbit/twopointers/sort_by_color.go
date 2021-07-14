package twopointers

func SortByColors(A []int) []int {
	var numOne, numTwo, numZero int
	for k := 0; k < len(A); k++ {
		if A[k] == 0 {
			numZero++
		} else if A[k] == 2 {
			numTwo++
		} else if A[k] == 1 {
			numOne++
		}
	}
	for i := 0; i < numZero; i++ {
		A[i] = 0
	}
	for i := numZero; i < numZero+numOne; i++ {
		A[i] = 1
	}
	for i := numZero + numOne; i < numOne+numTwo+numZero; i++ {
		A[i] = 2
	}

	return A
}
