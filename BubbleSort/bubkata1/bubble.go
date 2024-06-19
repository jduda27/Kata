package bubkata1

func BubbleSort(unsorted []int) []int {
	for i := 0; i < len(unsorted)-1; i++ {
		swapped := false
		for j := 0; j < len(unsorted)-i-1; j++ {
			if unsorted[j] > unsorted[j+1] {
				temp := unsorted[j+1]
				unsorted[j+1] = unsorted[j]
				unsorted[j] = temp
				swapped = true
			}

			if !swapped {
				break
			}
		}
	}
	return unsorted
}
