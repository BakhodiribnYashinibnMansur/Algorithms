package main

import (
	"math"
)

//
// Given a square matrix, calculate the absolute difference between the sums of its diagonals.
// For example, the square matrix  is shown below:
// 1 2 3
// 4 5 6
// 9 8 9
// The left-to-right diagonal =1+5+9=15 . The right to left diagonal =3+5+7 =17. Their absolute difference is|15-17|=2 .
//
func main() {
	diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}})
}

func diagonalDifference(arr [][]int32) int32 {
	var x1Total int32
	var x2Total int32
	for i := 0; i < len(arr); i++ {
		x1 := i
		x2 := len(arr) - 1 - i
		x1Total = x1Total + arr[i][x1]
		x2Total = x2Total + arr[i][x2]
	}
	return int32(math.Abs(float64(x2Total) - float64(x1Total)))
}
