package leetcode

func TrapRainWater(height []int) int {
	length := len(height)
	i := 0
	ans := 0
	maxLeft := 0
	for i := 0; i < length; i++ {
		if height[i] >= height[maxLeft] {
			maxLeft = i
		}
	}
	for j := 1; j < maxLeft; j++ {
		if height[j] < height[i] {
			ans += (height[i] - height[j])
		} else {
			i = j
		}
	}
	i = length - 1
	for j := length - 1; j >= maxLeft; j-- {
		if height[j] < height[i] {
			ans += (height[i] - height[j])
		} else {
			i = j
		}
	}
	return ans
}
