package arrays

import (
	"fmt"

	"hackersrank/util"
)

/**
Sample Input
4
1 4 3 2

Sample Output
2 3 4 1
*/
func ArraysDs() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	a, err := util.ScanSlice(n)
	if err != nil {
		panic(err)
	}
	loop := len(a) / 2
	last := len(a) - 1

	for i, num := range a {
		if i == loop {
			break
		}
		a[i] = a[last-i]
		a[last-i] = num
	}
	for _, num := range a {
		fmt.Printf("%d ", num)
	}
}
