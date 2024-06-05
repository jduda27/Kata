package bskata1

func Search(array []int, key int) (result int) {
	middle := len(array) / 2
	switch {
	case len(array) == 0:
		result = -1
	case array[middle] > key:
		result = Search(array[:middle], key)
	case array[middle] < key:
		result = Search(array[middle+1:], key)
		if result >= 0 {
			result += middle + 1
		}
	default:
		result = middle
	}
	return
}
