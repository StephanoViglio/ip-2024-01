package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numSeq[i])
	}
	if n > 0 {

		x := 1
		for i := 1; i < n; i++ {
			if numSeq[i] == numSeq[i-1] {
				x++
			}
		}

		fmt.Printf("A maior subsequencia de elementos iguais possui %d elementos", x)
	}

}
