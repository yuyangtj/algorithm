package sort

import (
	"errors"
	"sync"
)

var (
	ErrorNotImplemented = errors.New("Sorting failed")
)

//InsertionSort implements the inefficient insertion sort algorithm
func InsertionSort(s []int) ([]int, error) {
	length := len(s)
	for j := 1; j < length; j++ {
		key := s[j]
		i := j - 1
		for ; i >= 0 && s[i] > key; i-- {
			s[i+1] = s[i]
		}
		s[i+1] = key
	}
	return s, nil

}

// merge merges two ordered arrays into a new ordered array
func merge(A []int, p, q, r int) {
	L := append([]int(nil), A[p:q]...)
	R := append([]int(nil), A[q:r]...)
	i, j := 0, 0
	for k := p; i < len(L) && j < len(R) && k < r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		switch {
		case i == len(L):
			copy(A[k+1:], R[j:])
		case j == len(R):
			copy(A[k+1:], L[i:])
		}
	}

}

func mergeSort(A []int, p, r int) {
	if p < r-1 {
		q := (p + r) / 2
		mergeSort(A, p, q)
		mergeSort(A, q, r)
		merge(A, p, q, r)
	}
}

//MergeSort implements the inefficient merge sort algorithm
func MergeSort(A []int) ([]int, error) {
	mergeSort(A, 0, len(A))
	return A, nil
}

func mergeSortConcurr(A []int, p, r int) {
	if p < r-1 {
		q := (p + r) / 2
		var wg sync.WaitGroup
		wg.Add(2)
		go func(A []int, p, q int) {
			defer wg.Done()
			mergeSortConcurr(A, p, q)
		}(A, p, q)
		go func(A []int, q, r int) {
			defer wg.Done()
			mergeSortConcurr(A, q, r)
		}(A, q, r)
		wg.Wait()
		merge(A, p, q, r)
	}
}

func MergeSortConcurr(A []int) ([]int, error) {
	mergeSortConcurr(A, 0, len(A))
	return A, nil

}

type heap struct {
	array    []int
	length   int
	heapSize int
}

//func prtNum(i interface{}) {
//	r := reflect.ValueOf(i).String()
//	log.Printf("%T", r)
//}
func maxHeapify(h heap, index int) {
	l := index<<1 + 1

	r := index<<1 + 2

	var largest int

	if l < h.heapSize && h.array[l] > h.array[index] {

		largest = l
	} else {

		largest = index
	}
	if r < h.heapSize && h.array[r] > h.array[largest] {

		largest = r
	}
	if largest != index {

		h.array[index], h.array[largest] = h.array[largest], h.array[index]
		maxHeapify(h, largest)
	}
}

func buildMaxHeap(h heap) {
	h.heapSize = h.length
	index := h.length / 2

	for index >= 0 {

		maxHeapify(h, index)
		index--
	}
}

func heapSort(h heap) {
	buildMaxHeap(h)
	for i := h.length - 1; i >= 1; i-- {
		h.array[0], h.array[i] = h.array[i], h.array[0]
		h.heapSize--
		maxHeapify(h, 0)
	}
}

func HeapSort(a []int) ([]int, error) {
	length, heapSize := len(a), len(a)
	hp := heap{array: a, length: length, heapSize: heapSize}
	heapSort(hp)
	return a, nil
}
