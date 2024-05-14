package main

import (
	"fmt"
)

func main() {
	var n, vezesUltimo, maiorNota, indice, x int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	if n >= 1 && n <= 1000 {
		for i := 0; i < n; i++ {
			fmt.Scanln(&numSeq[i])

			if numSeq[i] < 0 || numSeq[i] > 10 {
				break
			}
		}

		for j := 0; j < n; j++ {
			if numSeq[j] == numSeq[n-1] {
				x++
				vezesUltimo = x
			}
		}

		for k := 0; k < n; k++ {
			if numSeq[k] > maiorNota {
				maiorNota = numSeq[k]
				indice = k
			}
		}

		fmt.Printf("Nota %d, %d vezes \n", numSeq[n-1], vezesUltimo)
		fmt.Printf("Nota %d, índice %d", maiorNota, indice)

	}

}
