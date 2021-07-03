package heap

type Heap struct {
	arr  []int
	size int
}

func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Left(i int) int {
	return (i * 2) + 1
}

func (h *Heap) Right(i int) int {
	return (i * 2) + 2
}

func (h *Heap) MaxHeapify(i int) {
	l := h.Left(i)
	r := h.Right(i)

	length := h.size
	var largest int
	if l < length && h.arr[l] > h.arr[i] {
		largest = l
	} else {
		largest = i
	}
	if r < length && h.arr[r] > h.arr[largest] {
		largest = r
	}
	if largest != i {
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		h.MaxHeapify(largest)
	}
}

func (h *Heap) BuildMaxHeap() {
	length := h.size
	for i := length/2 - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}

func (h *Heap) HeapSort() {
	h.BuildMaxHeap()
	for i := len(h.arr) - 1; i > 0; i-- {
		h.arr[0], h.arr[i] = h.arr[i], h.arr[0]
		h.size = h.size - 1
		h.MaxHeapify(0)
	}
}
