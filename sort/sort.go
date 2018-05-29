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
