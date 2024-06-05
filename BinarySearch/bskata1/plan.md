# Binary Search

Problem: `Given a sorted array of values, find the index of the desired key.`

Assumptions: The dataset is already sorted. We are using a small subset of possible values. We are using go
and do not know the size of the user's array input so we will use a slice instead.

Plan:

First we will create a method search that takes in the array, a slice of integers, and the desired key, a integer.

In this method we take the very middle index of the slice, if the value is equal to our key we return the value.
If the value of the slice is less than the middle index value, we recursively run the search method and take with the slice being from the 0 index to the middle index. If the value is greater we do the same but with the slice of the middle index to the length of the array instead.

If the length of the array is 1 and the key is not equal we return -1


