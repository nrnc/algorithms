package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m, n := len(nums1), len(nums2)

	var i, j, k, prev, ans int
	var even bool
	if (m+n)%2 == 0 {
		ans = m + (n-m)/2 + 1
		even = true
	} else {
		ans = m + (n-m+1)/2
	}
	k = 1
	for k = 1; k < ans; k++ {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				prev = nums1[i]
				i++
			} else {
				prev = nums2[j]
				j++
			}
		} else if i < m {
			prev = nums1[i]
			i++
		} else if j < n {
			prev = nums2[j]
			j++
		}
	}
	if k == ans {
		if !even {
			if i < m && j < n {
				if nums1[i] < nums2[j] {
					return float64(nums1[i])
				} else {
					return float64(nums2[j])
				}
			} else if i < m {
				return float64(nums1[i])
			} else if j < n {
				return float64(nums2[j])
			}
		} else {
			if i < m && j < n {
				if nums1[i] < nums2[j] {
					return float64(nums1[i]) + float64((prev-nums1[i]))/2.0
				} else {
					return float64(nums2[j]) + float64((prev-nums2[j]))/2.0
				}
			} else if i < m {
				return float64(nums1[i]) + float64((prev-nums1[i]))/2.0
			} else if j < n {
				return float64(nums2[j]) + float64((prev-nums2[j]))/2.0
			}
		}
	}
	return 0.0
}
