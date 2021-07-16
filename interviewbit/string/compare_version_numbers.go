package string

import (
	"strconv"
	"strings"
)

func CompareStrings(first, second string) int {
	first = strings.Trim(first, " ")
	second = strings.Trim(second, " ")

	first = strings.TrimLeft(first, "0")
	second = strings.TrimLeft(second, "0")

	if len(first) > len(second) {
		return 1
	} else if len(first) < len(second) {
		return -1
	} else {
		strLen := len(first)
		for i := 0; i < strLen; i++ {
			if string(first[i]) > string(second[i]) {
				return 1
			} else if string(first[i]) < string(second[i]) {
				return -1
			} else {
				continue
			}
		}
	}
	return 0
}

func CompareVersion(A string, B string) int {
	aSlice := strings.Split(A, ".")
	bSlice := strings.Split(B, ".")

	var i, j int

	for i < len(aSlice) && j < len(bSlice) {
		res := CompareStrings(aSlice[i], bSlice[j])
		if res == 1 {
			return 1
		} else if res == -1 {
			return -1
		}

		i++
		j++
	}

	for ; i < len(aSlice); i++ {
		first, _ := strconv.ParseInt(aSlice[i], 10, 64)
		if first == 0 {
			continue
		} else {
			return 1
		}
	}

	for ; j < len(bSlice); j++ {
		second, _ := strconv.ParseInt(bSlice[j], 10, 64)
		if second == 0 {
			continue
		} else {
			return -1
		}
	}

	if i == len(aSlice)-1 && j == len(bSlice)-1 {
		return 0
	}

	return 0

}
