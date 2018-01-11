package arrays

import (
	"fmt"
	"hackerrank/util"
	"strconv"
)

/**
Sample Input
5 4
1 2 3 4 5

Sample Output
5 1 2 3 4

When we perform  left rotations, the array undergoes the following sequence of changes:
[1,2,3,4,5] -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3,] -> [5,1,2,3,4]
*/
func Rotation() {
	var (
		n   int
		sft int
	)
	if _, err := fmt.Scan(&n); err != nil {
		panic(n)
	}
	if _, err := fmt.Scan(&sft); err != nil {
		panic(n)
	}

	a, err := util.ScanSlice(n)
	if err != nil {
		panic(err)
	}

	a = leftRotation(a, sft)
	//a = rightRotation(a, sft)
	for _, val := range a {
		fmt.Print(strconv.Itoa(val) + " ")
	}
}

func leftRotation(a []int, sft int) []int {
	a = append(a[sft:], a[:sft]...)
	return a
}

func rightRotation(a []int, sft int) []int {
	pos := len(a) - sft
	return append(a[pos:], a[0:pos]...)
}
