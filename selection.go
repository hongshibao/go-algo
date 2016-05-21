package algo

import (
	"errors"
	"sort"
)

func Partition(array sort.Interface, pivotIndex int) (int, error) {
	if pivotIndex < 0 || pivotIndex >= array.Len() {
		return -1, errors.New("Pivot index out of range")
	}
	lastIndex := array.Len() - 1
	array.Swap(pivotIndex, lastIndex)
	leftIndex := 0
	for i := 0; i < lastIndex; i++ {
		if array.Less(i, lastIndex) {
			array.Swap(leftIndex, i)
			leftIndex++
		}
	}
	array.Swap(leftIndex, lastIndex)
	return leftIndex, nil
}
