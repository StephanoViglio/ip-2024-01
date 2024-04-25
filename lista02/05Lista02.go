package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numSeq[i])
	}

	compSeq, compMax := 1, 0
	for i := 1; i < n; i++ {

		if numSeq[i] > numSeq[i-1] {
			compSeq++

			if compSeq > compMax {
				compMax = compSeq
			}
		}
	}
	fmt.Println("O comprimento do segmento crescente maximo e: ", compMax-1)
}
