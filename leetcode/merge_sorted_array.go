package leetcode

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	temp := make([]int, m)
	copy(temp, nums1)
	i := 0 // counter for temp array
	j := 0 // counter for nums2 array
	for k := 0; k < m+n; k++ {
		if i < m && j < n {
			if temp[i] < nums2[j] {
				nums1[k] = temp[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
		} else if i < m {
			nums1[k] = temp[i]
			i++
		} else if j < n {
			nums1[k] = nums2[j]
			j++
		}
	}
}
