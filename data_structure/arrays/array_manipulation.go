package arrays

import (
	"bufio"
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
sort.Sort 문제일까?
*/
func ArrayManipulation() {
	//var (
	//	listSize int
	//	n        int
	//)
	//if _, err := fmt.Scan(&listSize); err != nil {
	//	panic(err)
	//}
	//if _, err := fmt.Scan(&n); err != nil {
	//	panic(err)
	//}
	//
	//lines := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	a, err := util.IntScanlnSlice(n)
	//	if err != nil {
	//		panic(err)
	//	}
	//	lines[i] = a
	//}
	//
	//sequence := make([]int, listSize)
	//for _, item := range lines {
	//	updateSequenceByRangeIndex(sequence, item[0]-1, item[1]-1, item[2])
	//}
	//
	//a := sort.IntSlice(sequence)
	//sort.Sort(a)
	//fmt.Println(a[len(a)-1])
	//readWithScanner()
	readWithScanner2()
}

func updateSequenceByRangeIndex(sequence []int, start, end, val int) {
	for j := start; j <= end; j++ {
		sequence[j] += val
	}
}

func readWithScanner2() {
	var (
		listSize int
		n        int
	)
	file, err := os.Open("data_structure/arrays/input/array_manipulation11.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fscan(file, &listSize, &n); err != nil {
		panic(err)
	}
	fmt.Println("listSize", listSize, "n", n)

	lines := make([][]int, n)
	for i := 0; i < n; i++ {
		var (
			start int
			end   int
			val   int
		)

		if _, err := fmt.Fscanln(file, &start, &end, &val); err != nil {
			panic(err)
		}
		//a := [3]int{start, end, val}
		a := make([]int, n)
		a[0] = start
		a[1] = end
		a[2] = val

		//a := [n]int{
		//	start,
		//	end,
		//	val,
		//}

		lines[i] = a
	}
	sequence := make([]int, listSize)
	for _, item := range lines {
		updateSequenceByRangeIndex(sequence, item[0]-1, item[1]-1, item[2])
	}

	//a := sort.IntSlice(sequence)
	//sort.Sort(a)
	//fmt.Println(a[len(a)-1])
	fmt.Println(sequence[len(sequence)-1])
}

func readWithScanner() {
	// Open file and create scanner on top of it
	file, err := os.Open("data_structure/arrays/input/array_manipulation12.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanLines)
	//scanner.Split(bufio.ScanWords)

	for {
		// Scan for next token.
		success := scanner.Scan()
		if success == false {
			// False on error or EOF. Check error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
			} else {
				log.Fatal(err)
			}
			break
		}

		// Get data from scan with Bytes() or Text()

		//io.ByteReader()
		//strings.Split(scanner.Text(), " ")
		//fmt.Fscan()

		fmt.Println("First word found:", scanner.Text())
	}

	// Call scanner.Scan() again to find next token
}
