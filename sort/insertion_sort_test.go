package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{4, 7, 2, 8, 5, 0, 2}
	testArr := []int{0, 2, 2, 4, 5, 7, 8}
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, testArr) {
		t.Errorf("expected %v but got %v", testArr, arr)
	}
}

func TestInsertionSortDesc(t *testing.T) {
	arr := []int{4, 7, 2, 8, 5, 0, 2}
	testArr := []int{8, 7, 5, 4, 2, 2, 0}
	InsertionSortDesc(arr)
	if !reflect.DeepEqual(arr, testArr) {
		t.Errorf("expected %v but got %v", testArr, arr)
	}
}
