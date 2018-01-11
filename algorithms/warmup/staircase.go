package warmup

import "fmt"

func StaircasePrintMark(num int, mark string) {
	for i := 0; i < num; i++ {
		fmt.Print(mark)
	}
}

func Staircase() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	for i := 1; i <= n; i++ {
		StaircasePrintMark(n-i, " ")
		StaircasePrintMark(i, "#")
		StaircasePrintMark(1, "\n")
	}
}
