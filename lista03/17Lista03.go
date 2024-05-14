package main

import (
	"fmt"
)

func contadorElementosUnicos(numSeq []int) int {

	map_elementos := make(map[int]int)
	elementosUnicos := 0

	for _, i := range numSeq {
		map_elementos[i]++
	}
	for _, contador := range map_elementos {
		if contador == 1 {
			elementosUnicos++
		}
	}

	return elementosUnicos
}

func main() {

	var n int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	if n >= 1 && n <= 1000 {
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &numSeq[i])
		}

		quantElementosUnicos := contadorElementosUnicos(numSeq)

		fmt.Println(quantElementosUnicos)
	}
}
