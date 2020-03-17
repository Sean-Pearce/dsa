package mergesort

// MergeSort implements merge-sort algorithm
func MergeSort(arr []int) {
	mergesort(arr, 0, len(arr)-1)
}

func mergesort(arr []int, a, b int) {
	if b-a < 1 {
		return
	}

	m := (a + b) / 2
	mergesort(arr, a, m)
	mergesort(arr, m+1, b)
	merge(arr, a, m, b)
}

func merge(arr []int, a, m, b int) {
	if arr[m] <= arr[m+1] {
		return
	}

	tmp := make([]int, 0, b-a+1)
	i, j := a, m+1

	for i <= m && j <= b {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	for i <= m {
		tmp = append(tmp, arr[i])
		i++
	}
	for j <= b {
		tmp = append(tmp, arr[j])
		j++
	}

	i = a
	for _, v := range tmp {
		arr[i] = v
		i++
	}
}
