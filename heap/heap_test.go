package heap

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Heap{arr: []int{6, 3, 8, 4, 0, 2, 5, 7, 9, 9}, size: 10}
	h.HeapSort()
	testArr := []int{0, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	if !reflect.DeepEqual(testArr, h.arr) {
		t.Errorf("Expected %v but go %v\n", testArr, h.arr)
	}
}
