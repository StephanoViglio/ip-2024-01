package main

import (
	"fmt"
)

func selectionsort(numSeq []int) {
	var n = len(numSeq)

	for i := 0; i < n; i++ {
		var min = i
		for k := i; k < n; k++ {
			if numSeq[k] < numSeq[min] {
				min = k
			}
		}
		numSeq[i], numSeq[min] = numSeq[min], numSeq[i]
	}

}

func main() {

	var t int

	fmt.Scanln(&t)

	vector := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&vector[i])
	}

	fmt.Println("Não ordenado: \n", vector)

	selectionsort(vector)

	fmt.Println("Ordenado: \n", vector)
}
