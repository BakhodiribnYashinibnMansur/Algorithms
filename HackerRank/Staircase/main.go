package main

import (
	"fmt"
)

// Staircase detail

// This is a staircase of size n = 4:

//        #
//      ##
//   ###
// ####
// Its base and height are both equal to . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

// Write a program that prints a staircase of size .
func main() {
	staircase(4)
}

func staircase(n int32) {
	//My Solution

	// var i int32
	// var j int32
	// hash := ""
	// space := ""
	// var hashList []string
	// var spaceList []string
	// for i = 0; i < n; i++ {
	// 	hash = hash + "#"
	// 	hashList = append(hashList, hash)
	// }
	// for j = n - 1; j > 0; j-- {
	// 	space = space + " "
	// 	spaceList = append(spaceList, space)
	// }
	// for i = 0; i < n; i++ {
	// 	if i == n-1 {
	// 		fmt.Println(hashList[i])
	// 	} else {
	// 		fmt.Println(spaceList[n-2-i] + hashList[i])
	// 	}
	// }

	//HackerRank answer
	// var i int32 = 1
	// for ; i <= n; i++ {
	// 	fmt.Printf("%s%s\n",
	// 		strings.Repeat(" ", int(n-i)),
	// 		strings.Repeat("#", int(i)))
	// }

	// x := int(n)
	// hash := ""
	// space := strings.Repeat(" ", x-1)
	// res := ""

	// for i := 0; i < x; i++ {
	// 	hash = hash + "#"
	// 	res = space + hash
	// 	fmt.Println(res)
	// 	space = strings.TrimSuffix(space, " ")
	// }

	for i := 1; i <= int(n); i++ {
		for j := int(n) - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}
