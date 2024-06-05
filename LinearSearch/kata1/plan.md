# Linear Search

Problem: `We have an array A->N and we want to find if value x is in the array.`

Solution:
We have an array with a set amount of values in it. The size could be very small or very large and we want to find the if
the value x is in the array and what location in the array it is.
We assume size does not matter in this instance and we are only looking for the first occurrence from the left.


The way I would approach the problem would be to first create a method search that will take in both our array and the
value we are trying to find, key.

In this method we iterate over the length of the array and check if the value of the array at the current index is equal
to our desired key. If not we move forward to the next index in the array. The program returns the index of the first
instance of the key and exits the function at said occurence.

This method is very dependent on the size of the input array. It's time complexity is O(N) meaning in the worst case the
value is not found and it must check N values of the provided array.

Questions I might ask:
Are we wanting to find all locations of x in the array or just the first occurrence?
If it's only the first occurrence do we want it to be the first one from the left or right or middle?
How large of an array are we expecting to be searching over?
