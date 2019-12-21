package sort

// quick sort
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

// [3,5,1,5,2,4]
// i=0 j=0
// i=1 j=1
// i=1 j=2
// i=2 j=3 [3,1,5,5,2,4]
// i=2 j=4
// i=3 j=5 [3,1,2,5,5,4]
func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[j], arr[i] = arr[i], arr[j]
			}
			i++
		}
	}
	arr[i], arr[end] = pivot, arr[i]

	return i
}
