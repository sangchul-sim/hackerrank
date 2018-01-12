package util

import "fmt"

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func IntScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}

func IntScanlnSlice(n int) ([]int, error) {
	var intSlice []int
	for i := 0; i < n; i++ {
		var temp int
		if _, err := fmt.Scan(&temp); err != nil {
			return []int{}, err
		}
		intSlice = append(intSlice, temp)
	}
	return intSlice, nil
}

func StringScanlnSlice(n int) ([]string, error) {
	var stringSlice []string
	for i := 0; i < n; i++ {
		var temp string
		if _, err := fmt.Scan(&temp); err != nil {
			return []string{}, err
		}
		stringSlice = append(stringSlice, temp)
	}
	return stringSlice, nil
}
