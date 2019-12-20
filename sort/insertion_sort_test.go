package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 3, 1, 5, 4, 7, 6}
	InsertionSort(arr)
	t.Log(arr)
}
