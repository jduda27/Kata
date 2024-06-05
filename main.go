package main

import (
	"Kata/LinearSearch/kata1"
	"fmt"
)

func main() {
	array1 := []string{"A", "B", "C", "D", "E", "F", "H"}
	fmt.Println(array1)
	fmt.Printf("Linear Search\nKata 1 index of E: %d\n", kata1.Search(array1, "E"))
	fmt.Printf("Kata 1 index of G: %d\n", kata1.Search(array1, "G"))
}
