package main

// Given an integer x, return true if x is palindrome integer.
// An  integer is a palindrome when it reads the same backward as forward.
// For example, 121 is a palindrome while 123 is not.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
func main() {
	isPalindrome(121)
}
func isPalindrome(x int) bool {
	if x < 0 && x < 10 {
		return false
	}
	n := x
	reverseNumber := 0
	for n > 0 {
		remainder := n % 10
		reverseNumber = reverseNumber*10 + remainder
		n /= 10
	}
	if reverseNumber == x {
		return true
	}
	return false
}

// My Result
// Runtime: 8 ms, faster than 99.39% of Go online submissions for Palindrome Number.
// Memory Usage: 4.6 MB, less than 90.80% of Go online submissions for Palindrome Number.

//LeetCode Solution
// func isPalindrome(x int) bool {
// 	X := strconv.Itoa(x)
// 	l := len(X)
//     if x < 0 { return false }
//     for i := 0; i < len(X)/2; i++ {
// 		if X[i] == X[l-1] {
// 			l--
// 			continue
// 		}
// 		return false
// 	}
// 	return true
// }

// Runtime: 12 ms, faster than 84.19% of Go online submissions for Palindrome Number.
// Memory Usage: 5.2 MB, less than 62.50% of Go online submissions for Palindrome Number.
