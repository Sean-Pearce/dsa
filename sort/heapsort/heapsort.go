package heapsort

import (
	"container/heap"

	"github.com/Sean-Pearce/dsa/minheap"
)

func HeapSort(arr []int) []int {
	var h minheap.MinHeap = arr
	res := make([]int, 0, len(arr))
	heap.Init(&h)
	for h.Len() > 0 {
		res = append(res, heap.Pop(&h).(int))
	}
	return res
}
