package main

import (
	"fmt"
	"sort"
)

func iddPares(numSeq []int) {

	var numSeqPares = []int{}

	for i := 0; i < len(numSeq); i++ {
		if numSeq[i]%2 == 0 {
			numSeqPares = append(numSeqPares, numSeq[i])
		}
	}
	sort.Ints(numSeqPares)

	for k := 0; k < len(numSeqPares); k++ {
		fmt.Println(numSeqPares[k])
	}
}

func iddImpares(numSeq []int) {

	var numSeqImpares = []int{}

	for i := 0; i < len(numSeq); i++ {
		if numSeq[i]%2 != 0 {
			numSeqImpares = append(numSeqImpares, numSeq[i])
		}
	}
	sort.Ints(numSeqImpares)

	for k := len(numSeqImpares) - 1; k >= 0; k-- {
		fmt.Println(numSeqImpares[k])
	}
}

func main() {
	var n int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&numSeq[i])
	}

	iddPares(numSeq)
	iddImpares(numSeq)
}
