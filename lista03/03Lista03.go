package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var numSeq [5000]int

	if n < 5000 {
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &numSeq[i])
		}

		for i := n - 1; i >= 0; i-- {
			fmt.Print(numSeq[i], " ")
		}
	}
}
