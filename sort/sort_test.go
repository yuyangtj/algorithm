package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 6}, []int{1, 3, 4, 6}},
		{[]int{4, 1, 3}, []int{1, 3, 4}},
		{[]int{0}, []int{0}},
	}
	var testCases = []struct {
		name string
		item func([]int) ([]int, error)
	}{
		{"InsertionSort", InsertionSort},
		{"MergeSort", MergeSort},
		{"MergeSortConcurr", MergeSortConcurr},
		{"HeapSort", HeapSort},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for _, test := range tests {

				got, err := testCase.item(test.input)
				if err != nil {
					t.Fatal("sorting is not implemented")
				}
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("InsertionSort(%v) = %v, want %v", test.input, got, test.want)
				}
			}

		})

	}
}

func BenchmarkSort10k(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := r.Perm(10000)
	benchmarks := []struct {
		name string
		size int
		item func([]int) ([]int, error)
	}{
		{"InsertSort", len(l), InsertionSort},
		{"MergeSort", len(l), MergeSort},
		{"MergeSortConcurr", len(l), MergeSortConcurr},
		{"HeapSort", len(l), HeapSort},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.item(l)
			}

		})
	}
}
