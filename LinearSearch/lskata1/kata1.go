package lskata1

func Search(slice []string, key string) int {

	for i, v := range slice {
		if v == key {
			return i
		}
	}

	return -1

}
