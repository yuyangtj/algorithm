package sort

//InsertionSort implements the inefficient insertion sort algorithm
func InsertionSort(s []int) []int {
	length := len(s)
	for j := 1; j < length; j++ {
		key := s[j]
		i := j - 1
		for ; i >= 0 && s[i] > key; i-- {
			s[i+1] = s[i]
		}
		s[i+1] = key
	}
	return s

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

//MergeSort implements the inefficient merge sort algorithm
func MergeSort(A []int, p, r int) []int {
	if p < r-1 {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q, r)
		merge(A, p, q, r)
	}
	return A
}
