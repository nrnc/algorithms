package arrays

/**
 * @input A : Integer array
 *
 * @Output string.
 */
import (
	"sort"
	"strconv"
)

type IntArrayForStrings []int

func convItoa(i int) string {
	return strconv.Itoa(i)
}

func (r IntArrayForStrings) Less(i, j int) bool {
	return convItoa(r[i])+convItoa(r[j]) > convItoa(r[j])+convItoa(r[i])
}
func (r IntArrayForStrings) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r IntArrayForStrings) Len() int {
	return len(r)
}
func LargestNumber(A []int) string {
	sort.Sort(IntArrayForStrings(A))
	ans := ""
	for _, v := range A {
		if ans == "0" && convItoa(v) == "0" {
			continue
		}
		ans += convItoa(v)
	}

	return ans
}
