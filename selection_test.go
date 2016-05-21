package algo

import (
	"sort"
	"testing"
)

func checkPartitionResult(t *testing.T, array sort.Interface,
	idx int, ans int) {
	if idx != ans {
		t.Error("Partition result index is wrong")
	}
	for i := 0; i < idx; i++ {
		if !array.Less(i, idx) {
			t.Error("Partition result is wrong")
		}
	}
	for i := idx + 1; i < array.Len(); i++ {
		if !array.Less(idx, i) {
			t.Error("Partition result is wrong")
		}
	}
}

func TestPartition(t *testing.T) {
	// case 1
	array := sort.IntSlice{2, 3, 4, 1, 5}
	idx, err := Partition(array, 1)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, idx, 2)
	// case 2
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 5)
	if err == nil {
		t.Error("Partition should return error here")
	}
	// case 3
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 4)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, idx, 4)
	// case 4
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 3)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, idx, 0)
}
