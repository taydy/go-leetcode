package sort

// 插入排序
func InsertionSort(arr []int) {
	length := len(arr)

	if length == 0 {
		return
	}

	for i := 1; i < length; i++ {

		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] <= value {
				break
			}

			arr[j+1] = arr[j]
		}
		arr[j+1] = value
	}
}
