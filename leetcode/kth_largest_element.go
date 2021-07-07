package leetcode

import "container/heap"

type IntArray []int

func (h IntArray) Less(i, j int) bool {
	return h[i] >= h[j]
}
func (h IntArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntArray) Len() int {
	return len(h)
}
func (h *IntArray) Pop() interface{} {
	tmp := *h
	n := len(tmp)
	var res interface{} = tmp[n-1]
	*h = tmp[:n-1]
	return res
}
func (h *IntArray) Push(num interface{}) {
	*h = append(*h, num.(int))
}
func FindKthLargest(nums []int, k int) int {
	h := &IntArray{}
	for _, v := range nums {
		*h = append(*h, v)
	}
	heap.Init(h)
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}

	return res[k-1]
}
