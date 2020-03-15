package minheap

// MinHeap is min-heap of ints
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds a num to h
func (h *MinHeap) Push(x interface{}) {
	(*h) = append(*h, x.(int))
}

// Pop removes the minimum value from h and returns it
func (h *MinHeap) Pop() interface{} {
	l := len(*h)
	last := (*h)[l-1]
	*h = (*h)[:l-1]
	return last
}
