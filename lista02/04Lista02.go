package main

import (
	"fmt"
)

func main() {
	var k int
	var n, i, s float64

	fmt.Scanln(&n) //número a ser utilizado
	fmt.Scanln(&i) //valor incial da soma antes de ter incremento
	fmt.Scanln(&k) //K vezes
	fmt.Scanln(&s) //incremento que se repetirá

	incremento := make([]float64, k)
	incrementos := make([]float64, k)
	soma := make([]float64, k)
	subtracao := make([]float64, k)
	multiplicacao := make([]float64, k)
	divisao := make([]float64, k)

	if n >= 0 && n <= 9 {

		fmt.Println("Tabuada de soma: ")

		var y float64 = 0
		for x := 0; x < k; x++ {

			incremento[x] = s * y
			incrementos[x] = i + incremento[x]
			soma[x] = n + i + incremento[x]
			fmt.Println(n, " + ", incrementos[x], " = ", soma[x])
			y++
		}

		fmt.Println("Tabuada de substração: ")

		var w float64 = 0
		for x := 0; x < k; x++ {

			incremento[x] = s * w
			incrementos[x] = i + incremento[x]
			subtracao[x] = n - i - incremento[x]
			fmt.Println(n, " - ", incrementos[x], " = ", subtracao[x])
			w++
		}

		fmt.Println("Tabuada de multiplicacao: ")

		var z float64 = 0
		for x := 0; x < k; x++ {

			incremento[x] = s * z
			incrementos[x] = i + incremento[x]
			multiplicacao[x] = n * incrementos[x]
			fmt.Println(n, " x ", incrementos[x], " = ", (multiplicacao[x] * 100 / (100)))
			z++
		}

		fmt.Println("Tabuada de divisão: ")

		var q float64 = 0
		for x := 0; x < k; x++ {

			incremento[x] = s * q
			incrementos[x] = i + incremento[x]
			divisao[x] = (n) / (incrementos[x])
			fmt.Println(n, " / ", incrementos[x], " = ", (divisao[x] * 100 / (100)))
			q++
		}

	}

}
