package main

import (
	"fmt"
	"sort"
)

func detMenorDif(vetor []float64) float64 {
	var menorDif float64
	menorDif = 10000

	sort.Float64s(vetor)

	for i := 1; i < len(vetor); i++ {

		dif := (vetor[i] - vetor[i-1])
		if dif < menorDif {
			menorDif = dif
		}

	}
	return menorDif
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	factorialOfN := n * factorial(n-1)

	return factorialOfN
}

func combinacaoDeNpor2(n int) int {
	numerador := factorial(n)
	denominadorFact := factorial(n - 2)
	denominador := 2 * denominadorFact

	combinacao := (numerador) / (denominador)

	return combinacao
}

func main() {
	var t, n int

	fmt.Scanln(&t)

	var difCadaVetor [99]float64
	var combinacaoCadaN [99]int

	if t >= 1 && t <= 10 {
		for k := 0; k < t; k++ {
			fmt.Scanln(&n)

			vetor := make([]float64, n)

			if n >= 2 && n <= 1000 {
				for i := 0; i < n; i++ {

					fmt.Scanln(&vetor[i])
				}

				difCadaVetor[k] = detMenorDif(vetor)
				vetor = nil
				combinacaoCadaN[k] = combinacaoDeNpor2(n)

			}

		}
	}

	for i := 0; i < t; i++ {
		fmt.Println(difCadaVetor[i], " ", combinacaoCadaN[i])
	}

}
