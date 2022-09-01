// 13. Roman to Integer

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.

// Example 1:

// Input: s = "III"
// Output: 3
// Explanation: III = 3.
// Example 2:

// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
// Example 3:

// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

package main

import "fmt"

func main() {
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	romanNumberMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	RomanToInt := romanNumberMap[string(s[len(s)-1])]
	for i := range s {
		if i < len(s)-1 {
			if romanNumberMap[string(s[i])] < romanNumberMap[string(s[i+1])] {
				RomanToInt -= romanNumberMap[string(s[i])]
			} else {
				RomanToInt += romanNumberMap[string(s[i])]
			}
		}
	}
	fmt.Println(RomanToInt)
	return RomanToInt
}

// Runtime: 6 ms, faster than 88.29% of Go online submissions for Roman to Integer.
// Memory Usage: 2.9 MB, less than 29.41% of Go online submissions for Roman to Integer.

// func romanToInt(s string) int {
// 	data := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
// 	ret := data[s[len(s)-1]]
// 	for i := len(s) - 2; i >= 0; i-- {
// 		if data[s[i]] < data[s[i+1]] {
// 			ret -= data[s[i]]
// 		} else {
// 			ret += data[s[i]]
// 		}
// 	}
// 	return ret
// }
// Runtime: 7 ms, faster than 87.48% of Go online submissions for Roman to Integer.
// Memory Usage: 2.9 MB, less than 81.32% of Go online submissions for Roman to Integer.
