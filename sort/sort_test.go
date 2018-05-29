package sort

import "testing"

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}
func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 6}, []int{1, 3, 4, 6}},
		{[]int{4, 1, 3}, []int{1, 3, 4}},
	}

	for _, test := range tests {
		if got := InsertionSort(test.input); !compareSlice(got, test.want) {
			t.Errorf("InsertionSort(%v) = %v", test.input, got)
		}
	}
}
