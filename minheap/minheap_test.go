package minheap

import (
	"container/heap"
	"reflect"
	"testing"
)

var tests = []struct {
	input []int
	want  []int
}{
	{[]int{3, 2, 1}, []int{1, 2, 3}},
}

func TestMinHeap(t *testing.T) {
	for _, test := range tests {
		hh := append(MinHeap{}, test.input...)
		h := &hh
		heap.Init(h)
		got := make([]int, 0, h.Len())
		for h.Len() > 0 {
			got = append(got, heap.Pop(h).(int))
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want: %v, got: %v", test.want, got)
		}
	}
}
