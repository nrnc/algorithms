package leetcode

func LengthOfLongestSubstring(s string) int {
	indexes := make(map[byte]int)
	left, right, ans := 0, 0, 0
	for right < len(s) {
		if index, ok := indexes[s[right]]; ok && index >= left {
			left = index + 1
		}
		indexes[s[right]] = right
		right++
		ans = max(ans, right-left)
	}

	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
