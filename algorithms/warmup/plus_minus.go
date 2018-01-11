package warmup

import (
	"fmt"

	"hackerrank/util"
)

func PlusMinus() {
	var (
		n   int
		err error
	)

	_, err = fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	a, err := util.IntScanln(n)
	if err != nil {
		panic(err)
	}
	var (
		positive int
		negative int
		zero     int
	)

	for _, num := range a {
		if num > 0 {
			positive++
		} else if num == 0 {
			zero++
		} else if num < 0 {
			negative++
		}
	}

	fmt.Printf("%f\n", float64(positive)/float64(n))
	fmt.Printf("%f\n", float64(negative)/float64(n))
	fmt.Printf("%f\n", float64(zero)/float64(n))
}
