package main

import (
	"fmt"
	"sort"
)

func mediana(l int, numSeq []float64) float64 {

	var media float64
	var x int

	if l%2 == 0 {
		for j := 1; j <= l-1; j++ {
			x = l / 2
			media = (numSeq[x] + numSeq[x-1]) / (2)
		}
	} else {

		for j := 0; j < l; j++ {
			x = l / 2
			media = (numSeq[x])
		}

	}
	return media
}

func main() {
	var n int

	fmt.Scanln(&n)

	numSeq := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&numSeq[i])
	}

	sort.Float64s(numSeq)

	fmt.Println(mediana(n, numSeq))

}
