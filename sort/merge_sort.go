package sort

type ints []int

func (i ints) Iter() func() (int, bool) {
	index := 0
	return func() (val int, ok bool) {
		if index >= len(i) {
			return
		}
		val, ok = i[index], true
		index++
		return
	}
}

func Merge(arr []int) {
	i := len(arr) / 2
	leftArray := make([]int, i)
	rightArray := make([]int, len(arr)-i)
	copy(leftArray, arr[:i])
	copy(rightArray, arr[i:])
	leftIter, rightIter := ints(leftArray).Iter(), ints(rightArray).Iter()
	leftValue, leftHasNext := leftIter()
	rightValue, rightHasNext := rightIter()
	for k := range arr {
		if !leftHasNext {
			arr[k] = rightValue
			rightValue, rightHasNext = rightIter()
		} else if !rightHasNext {
			arr[k] = leftValue
			leftValue, leftHasNext = leftIter()
		} else {
			if leftValue < rightValue {
				arr[k] = leftValue
				leftValue, leftHasNext = leftIter()
			} else {
				arr[k] = rightValue
				rightValue, rightHasNext = rightIter()
			}
		}
	}
}

func MergeSort(arr []int) {
	i := len(arr) / 2
	if i > 0 {
		MergeSort(arr[:i])
		MergeSort(arr[i:])
		Merge(arr)
	}
}
