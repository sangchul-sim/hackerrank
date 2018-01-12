package arrays

import (
	"fmt"
	"sort"

	"hackerrank/util"
)

/**
hourglass
a b c
  d
e f g

Sample Input
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0

Sample Output
19

Sample Input
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 9 2 -4 -4 0
0 0 0 -2 0 0
0 0 -1 -2 -4 0

Sample Output
13
*/
func TwoDArrayDs() {
	var (
		n     = 6 // n by n matrix
		width = 3 // n by n sub matrix
	)
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		a, err := util.IntScanlnSlice(n)
		if err != nil {
			panic(err)
		}
		matrix[i] = a
	}

	subMatrix := [][]int{}
	for i, _ := range matrix {
		if len(matrix[i:]) < width {
			continue
		}
		for _, mtx := range makeSubMatrix(matrix[i:i+width], width) {
			subMatrix = append(subMatrix, mtx)
		}
	}

	sumSubMatrix := make([]int, len(subMatrix))
	for i, mtx := range subMatrix {
		sumSubMatrix[i] = sumHourGlass(mtx)
	}

	a := sort.IntSlice(sumSubMatrix[0:])
	sort.Sort(a)
	fmt.Println(a[len(a)-1])
}

/**
a b c
  e
g h i
*/
func sumHourGlass(data []int) int {
	var result int
	for i, val := range data {
		switch i {
		case 3, 5:
			continue
		default:
			result += val
		}
	}
	return result
}

func makeSubMatrix(data [][]int, width int) map[int][]int {
	matrixMap := map[int][]int{}
	for _, sub := range data {
		var pos int
		for j, _ := range sub {
			if len(sub[j:]) >= width {
				if _, ok := matrixMap[pos]; !ok {
					matrixMap[pos] = []int{}
				}
				matrixMap[pos] = append(matrixMap[pos], sub[j:j+width]...)
				pos++
			}
		}
	}
	return matrixMap
}
