package arrays

import (
	"fmt"
	"log"
	"os"
)

/**
You are given a list(1-indexed) of size n, initialized with zeroes.
You have to perform m operations on the list and output the maximum of final values of all the n elements in the list.
For every operation, you are given three integers a, b and k and you have to add value k to all the elements ranging
from index a to b(both inclusive).

For example, consider a list a of size 3. The initial list would be a = [0, 0, 0]
and after performing the update 0(a,b,k) = (2, 3, 30), the new list would be a = [0, 30, 30].
Here, we've added value 30 to elements between indices 2 and 3. Note the index of the list starts from 1.

Input Format

The first line will contain two integers n and m separated by a single space.
Next m lines will contain three integers a, b and k separated by a single space.
Numbers in list are numbered from 1 to n.

Sample Input
5 3
1 2 100
2 5 100
3 4 100

Sample Output
200

Explanation
After first update list will be 100 100 0 0 0.
After second update list will be 100 200 100 100 100.
After third update list will be 100 200 200 200 100.
So the required answer will be 200.

시간이 오래 걸려서 test case failed..

*/
func ArrayManipulation() {
	var (
		n int
		m int
	)
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	if _, err := fmt.Scan(&m); err != nil {
		panic(err)
	}

	sequence := make([]int, n)
	for i := 0; i < m; i++ {
		var (
			a int
			b int
			k int
		)

		if _, err := fmt.Scanln(&a, &b, &k); err != nil {
			panic(err)
		}
		//updateSequenceByRangeIndex(sequence, a-1, b-1, k)
		//
		sequence[a-1] += k
		if b < n {
			sequence[b] -= k
		}
	}

	//a := sort.IntSlice(sequence)
	//sort.Sort(a)
	//fmt.Println(a[len(a)-1])
	//
	sum := 0
	max := 0
	for i := 0; i < n; i++ {
		sum += sequence[i]
		if max < sum {
			max = sum
		}
	}
	fmt.Println(max)
}

func updateSequenceByRangeIndex(sequence []int, start, end, val int) {
	for j := start; j <= end; j++ {
		sequence[j] += val
	}
}

func ArrayManipulationCase11() {
	var (
		n int
		m int
	)
	file, err := os.Open("data_structure/arrays/input/array_manipulation11.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fscan(file, &n, &m); err != nil {
		panic(err)
	}

	sequence := make([]int, n)
	for i := 0; i < m; i++ {
		var (
			a int
			b int
			k int
		)
		if _, err := fmt.Fscanln(file, &a, &b, &k); err != nil {
			panic(err)
		}
		sequence[a-1] += k
		if b < n {
			sequence[b] -= k
		}
	}

	sum := 0
	max := 0
	for i := 0; i < n; i++ {
		sum += sequence[i]
		if max < sum {
			max = sum
		}
	}
	fmt.Println(max)
}
