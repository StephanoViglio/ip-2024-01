package main

import (
	"fmt"
)

func troca(vetor []int, i, j int) {

	vetor[j], vetor[i] = vetor[i], vetor[j]
}

func trocaOpostoSeMenor(vetor []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			termoOposto := n - (i + 1)
			if j > i && j == termoOposto && vetor[i] < vetor[j] {
				troca(vetor, i, j)
			}
		}
	}
}

func main() {
	var t, n int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&n)
		if n < 500 && n > 0 {
			vetor := make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scanln(&vetor[j])

			}
			trocaOpostoSeMenor(vetor, n)
			fmt.Println(vetor)
		}
	}
}
