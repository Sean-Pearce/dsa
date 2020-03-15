package heapsort

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input []int
	want  []int
}{
	{[]int{}, []int{}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{5, 2, 20, 7, 3, 6, 17, 1, 9, 4, 14, 15, 13, 11, 19, 10, 12, 18, 16, 8},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
}

func TestHeapSort(t *testing.T) {
	for _, test := range tests {
		if got := HeapSort(append([]int{}, test.input...)); !reflect.DeepEqual(got, test.want) {
			t.Errorf("HeapSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(append([]int{}, tests[2].input...))
	}
}