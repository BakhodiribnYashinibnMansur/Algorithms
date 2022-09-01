package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.
// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

func main() {
	twoSum([]int{2, 7, 11, 15, 5, 4}, 9)
}

func twoSum(nums []int, target int) []int {
	var result []int
	for i, numi := range nums {
		for j, numj := range nums {
			if i != j && numi+numj == target {
				result = append(result, i)
			}
		}
	}
	fmt.Println("Result", result)
	return result
}

// My Result
// Runtime: 159 ms, faster than 5.05% of Go online submissions for Two Sum.
// Memory Usage: 3.7 MB, less than 67.09% of Go online submissions for Two Sum.
// Next challenges:

// LeetCodeSolution
// Runtime: 4 ms, faster than 97.78% of Go online submissions for Two Sum.
// Memory Usage: 4.3 MB, less than 25.57% of Go online submissions for Two Sum.

// func twoSum(nums []int, target int) []int {
//     tmp := make(map[int]int)
//     tmp[nums[0]] = 0
//     for i := 1; i < len(nums); i++ {
//         if j, ok := tmp[target - nums[i]]; ok { return []int{i, j} }
//         tmp[nums[i]] = i
//     }
//     return []int{}
// }
