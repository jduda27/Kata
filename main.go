package main

import (
	"Kata/BinarySearch/bskata1"
	"Kata/LinearSearch/lskata1"
	"fmt"
)

func main() {
	array1 := []string{"A", "B", "C", "D", "E", "F", "H"}
	fmt.Println(array1)
	fmt.Printf("Linear Search\nKata 1 index of E: %d\n", lskata1.Search(array1, "E"))
	fmt.Printf("Kata 1 index of G: %d\n", lskata1.Search(array1, "G"))

	array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array2)
	fmt.Printf("Binary Search\nKata 2 index of 3: %d\n", bskata1.Search(array2, 3))
	fmt.Printf("Kata 2 index of 20: %d\n", bskata1.Search(array2, 20))
	fmt.Printf("Kata 2 index of 10: %d\n", bskata1.Search(array2, 10))

}
