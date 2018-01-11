package warmup

import (
	"fmt"
	"sort"

	"hackersrank/util"
)

func MiniMaxSum() {
	n := 5
	sumArr := make([]int, n)

	a, err := util.IntScanln(5)
	if err != nil {
		panic(err)
	}

	sumExceptSome := func(data []int, idx int) int {
		var result int
		for i, num := range data {
			if i != idx {
				result += num
			}
		}
		return result
	}
	for i, _ := range a {
		sumArr[i] = sumExceptSome(a, i)
	}
	sort.Sort(util.IntSlice(sumArr))
	fmt.Println(sumArr[0], sumArr[len(sumArr)-1])
}
