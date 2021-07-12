package math

func ExcelColumNumber(A string) int {
	var res, val int
	for _, v := range A {
		val = int(v - 'A' + 1)
		res = res*26 + val
	}

	return res
}
