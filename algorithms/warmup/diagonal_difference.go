package warmup

import (
	"fmt"
	"math"

	"hackersrank/util"
)

func intSum(data []int) int {
	var result int
	for _, num := range data {
		result += num
	}
	return result
}

func DiagonalDifference() {
	var n int
	var err error
	fmt.Scan(&n)

	sq := make([][]int, n)
	for i := 0; i < n; i++ {
		sq[i], err = util.IntScanln(n)
	}

	if err != nil {
		panic(err)
	}

	pDiagonal := make([]int, n)
	sDiagonal := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				pDiagonal = append(pDiagonal, sq[i][j])
			}
			if i+j == n-1 {
				sDiagonal = append(sDiagonal, sq[i][j])
			}
		}
	}
	sumPrimaryDiagonal := intSum(pDiagonal)
	sumSecondaryDiagonal := intSum(sDiagonal)
	fmt.Println(math.Abs(float64(sumPrimaryDiagonal - sumSecondaryDiagonal)))
}
