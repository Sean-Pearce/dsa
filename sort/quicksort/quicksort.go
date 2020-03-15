package quicksort

// QuickSort implements quick-sort algorithm
func QuickSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	pivot := partition(arr, 0, len(arr)-1)
	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
	return arr
}

func partition(arr []int, low, high int) int {
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < arr[high] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
