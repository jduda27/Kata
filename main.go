package main

import (
	"Kata/BinarySearch/bskata1"
	"Kata/BubbleSort/bubkata1"
	"Kata/Excersizes/TwoCrystalBalls"
	"Kata/LinearSearch/lskata1"
	"Kata/LinkedList/dlkata1"
	"Kata/LinkedList/llkata1"
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
	fmt.Printf("Kata 2 index of 1: %d\n", bskata1.Search(array2, 1))

	building := []bool{false, false, false, false, false, true, true, true, true, true}
	fmt.Println(building)
	fmt.Printf("Two Crystal Balls\nOptimal Floor to Break Crystal Balls: %d\n", tcb.LowestBreakingPoint(building))

	array3 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(array3)
	fmt.Printf("BubbleSort\nKata 1 Sorted:\n%d\n", bubkata1.BubbleSort(array3))

	s1 := llkata1.List{}
	s1.Insert(1)
	s1.Insert(2)
	s1.Insert(3)

	llkata1.Show(&s1)

	s2 := dlkata1.List{}
	s2.Insert(4)
	s2.Insert(5)
	s2.Insert(6)
	s2.Insert(7)

	dlkata1.Show(&s2)

}
