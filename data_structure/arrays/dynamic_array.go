package arrays

import (
	"fmt"
	"hackerrank/util"
)

/**
n : number of sequence
lastAns : 0
q : number of query sequences S0, S1, ... Sn

for queries:
type(1) : 1 x y
	index of sequence id = (x ^ lastAns) % n
	append y in Si-th sequence

type(2) : 2 x y
	index of sequence id = (x ^ lastAns) % n
	lastAns = valueAt(y % (size of Si-th sequence))
	and print lastAns

Sample Input
2 5
1 0 5
1 1 7
1 0 3
2 1 0
2 1 1

Sample Output
7
3
*/
func DynamicArray() {
	const (
		queryType1    = 1
		queryType2    = 2
		queryParamCnt = 3
	)
	var (
		n        int
		queryCnt int
		lastAns  int
	)
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	if _, err := fmt.Scan(&queryCnt); err != nil {
		panic(err)
	}

	queries := make([][]int, queryCnt)
	sequence := make([][]int, n)

	for i, _ := range queries {
		data, err := util.IntScanln(queryParamCnt)
		if err != nil {
			panic(err)
		}
		queries[i] = append(queries[i], data...)
	}

	for _, query := range queries {
		queryType := query[0]
		x := query[1]
		y := query[2]
		seqId := getSeqId(n, x, lastAns)

		switch queryType {
		case queryType1:
			sequence[seqId] = append(sequence[seqId], y)
		case queryType2:
			pos := y % len(sequence[seqId])
			lastAns = sequence[seqId][pos]
			fmt.Println(lastAns)
		}
	}
}

func getSeqId(n, x, lastAns int) int {
	return (x ^ lastAns) % n
}
