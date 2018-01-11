package warmup

import (
	"fmt"
	"time"
)

func TimeConversion() {
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		panic(err)
	}
	t, err := time.Parse("3:04:05PM", input)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format("15:04:05"))
}
