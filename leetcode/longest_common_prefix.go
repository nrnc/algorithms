package leetcode

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLength := 201 // as given in PS that len(str) <= 200
	first := strs[0]
	for _, v := range strs {
		if len(v) < minLength {
			minLength = len(v)
			first = v
		}
	}
	prefix := make([]byte, 0, minLength)
	// outerloop on minLength string
	for i, v := range []byte(first) {
		j := 0
		// inner loop on all the strings
		for ; j < len(strs); j++ {
			if strs[j][i] != (v) {
				break
			}
		}
		if j == len(strs) {
			prefix = append(prefix, (v))
		} else {
			break
		}
	}

	return string(prefix)
}
