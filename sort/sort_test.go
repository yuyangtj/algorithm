package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
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
			t.Errorf("InsertionSort(%v) = %v, want %v", test.input, got, test.want)
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
			t.Errorf("InsertionSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestMergeSortConcurr(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 6}, []int{1, 3, 4, 6}},
		{[]int{4, 1, 3}, []int{1, 3, 4}},
	}

	for _, test := range tests {
		if got := MergeSortConcurr(test.input, 0, len(test.input)); !reflect.DeepEqual(got, test.want) {
			t.Errorf("InsertionSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := r.Perm(10000)
	for i := 0; i < b.N; i++ {
		InsertionSort(l)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := r.Perm(10000)
	for i := 0; i < b.N; i++ {
		MergeSort(l, 0, len(l))
	}
}

func BenchmarkMergeSortConcurr(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := r.Perm(10000)
	for i := 0; i < b.N; i++ {
		MergeSortConcurr(l, 0, len(l))
	}
}
