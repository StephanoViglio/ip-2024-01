package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	if n > 1 && n <= 1000 {
		for i := 0; i < n; i++ {
			fmt.Scanln(&numSeq[i])
		}

		sort.Ints(numSeq)

		for i := 0; i < n; i++ {
			fmt.Println(numSeq[i])
		}

	}
}
