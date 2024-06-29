package main

import (
	"fmt"
)

func ordenar(vetor []int) {
	l := len(vetor)
	for i := 0; i < l-1; i++ {
		for j := l - 1; j > i; j-- {
			if vetor[j-1] > vetor[j] {
				vetor[j-1], vetor[j] = vetor[j], vetor[j-1]
			}
		}
	}
}

func main() {
	vetorExemplo := []int{9, 8, 7, 6, 5, 4}
	ordenar(vetorExemplo)
	fmt.Println(vetorExemplo)
}
