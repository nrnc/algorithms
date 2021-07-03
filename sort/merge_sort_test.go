package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{4, 7, 2, 8, 5, 0, 2}
	testArr := []int{0, 2, 2, 4, 5, 7, 8}

	MergeSort(arr)

	if !reflect.DeepEqual(arr, testArr) {
		t.Errorf("expected %v got %v", testArr, arr)
	}
}
