package mergesort

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

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		got := append([]int{}, test.input...)
		if MergeSort(got); !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
