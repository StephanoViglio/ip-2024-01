package main

import (
	"fmt"
)

func main() {
	var n, m, x int

	var maiorElemento [999]int
	var maiorIndice [999]int

	m = 1
	n = 2

	if n > 1 && n <= 10000 {
		for i := 0; m != 0; i++ {
			fmt.Scanln(&n)
			m = n
			vetor := make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scanln(&vetor[j])
				if m == 0 {
					break
				}
			}
			for k := 0; k < n; k++ {
				if vetor[k] > maiorElemento[i] {
					maiorElemento[i] = vetor[k]
					maiorIndice[i] = k
				}
			}
			x++
			vetor = vetor[:0]
		}
		for i := 0; i < x-1; i++ {
			fmt.Println(maiorIndice[i], " ", maiorElemento[i])

		}

	}

}
