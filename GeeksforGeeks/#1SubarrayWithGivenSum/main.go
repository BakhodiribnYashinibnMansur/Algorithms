package main

import "fmt"

// Subarray with given sum
// EasyAccuracy: 39.71%Submissions: 100k+Points: 2
// Lamp
// This problem is part of GFG SDE Sheet. Click here to view more.

// Given an unsorted array A of size N that contains only non-negative integers, find a continuous sub-array which adds to a given number S.

// In case of multiple subarrays, return the subarray which comes first on moving from left to right.

// Example 1:

// Input:
// N = 5, S = 12
// A[] = {1,2,3,7,5}
// Output: 2 4
// Explanation: The sum of elements
// from 2nd position to 4th position
// is 12.

// Example 2:

// Input:
// N = 10, S = 15
// A[] = {1,2,3,4,5,6,7,8,9,10}
// Output: 1 5
// Explanation: The sum of elements
// from 1st position to 5th position
// is 15.
func main() {
	subarraySum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 14)
}

func subarraySum(numbers []int, target int) []int {
	var sum int
	start := -1
	end := -1
	for i := 0; i < len(numbers); i++ {
		sum = 0
		for j := i; j < len(numbers); j++ {
			sum = sum + numbers[j]
			if sum == target {
				start = i
				end = j
			}
		}
	}
	fmt.Println("start: ", start, " end: ", end)
	return []int{start, end}
}

