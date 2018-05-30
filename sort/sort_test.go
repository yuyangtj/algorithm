package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 6}, []int{1, 3, 4, 6}},
		{[]int{4, 1, 3}, []int{1, 3, 4}},
	}

	for _, test := range tests {
		if got := InsertionSort(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("InsertionSort(%v) = %v", test.input, got)
		}
	}
}

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 6}, []int{1, 3, 4, 6}},
		{[]int{4, 1, 3}, []int{1, 3, 4}},
	}

	for _, test := range tests {
		if got := MergeSort(test.input, 0, len(test.input)); !reflect.DeepEqual(got, test.want) {
			t.Errorf("InsertionSort(%v) = %v", test.input, got)
		}
	}
}
