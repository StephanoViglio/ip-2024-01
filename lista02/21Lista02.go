package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanln(&n)

	var numSeq [999]float64

	for i := 0; i < n; i++ {

		fmt.Scan(&numSeq[i])

		if n == 0 {
			break
		}
	}

	for i := 1; i <= n; i++ {

		if numSeq[i] > numSeq[i-1] {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}

	}
}
