package warmup

import (
	"fmt"
	"sort"

	"hackerrank/util"
)

func BirthdayCakeCandles() {
	var (
		n   int
		err error
	)
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	a, err := util.IntScanln(n)
	if err != nil {
		panic(err)
	}

	candleMap := make(map[int]int)

	for _, num := range a {
		candleMap[num]++
	}
	sort.Sort(util.IntSlice(a))
	tallestCandleHeight := a[len(a)-1]
	if count, ok := candleMap[tallestCandleHeight]; ok {
		fmt.Println(count)
	}
}
