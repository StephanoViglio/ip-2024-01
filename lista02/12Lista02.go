package main

import (
	"fmt"
)

func main() {

	var valorIngresso, valorInicial, valorFinal, maiorLucro, codML, maiorNumIngresso int

	fmt.Scanln(&valorIngresso)
	fmt.Scanln(&valorInicial)
	fmt.Scanln(&valorFinal)

	var valorV, dif, n, lucro, receita, despesas [9999]int

	if valorInicial < valorFinal {

		w := valorInicial
		for i := valorInicial; i <= valorIngresso; i++ {

			valorV[i] += i
			dif[i] = (valorIngresso - (valorV[i])) * 50
			n[i] = dif[i] + 120
			receita[i] = valorV[i] * n[i]
			despesas[i] = 200 + n[i]/(20)
			lucro[i] = receita[i] - despesas[i]

			fmt.Println("V: ", w, ", N: ", n[i], ", L: ", lucro[i])
			w++

			if lucro[i] > maiorLucro {
				maiorLucro = lucro[i]
				maiorNumIngresso = n[i]
				codML = w - 1
			}

		}

		z := valorIngresso + 1
		for i := valorIngresso + 1; i <= valorFinal; i++ {

			valorV[i] += i
			dif[i] = ((valorV[i]) - valorIngresso) * 60
			n[i] = 120 - dif[i]
			receita[i] = valorV[i] * n[i]
			despesas[i] = 200 + n[i]/(20)
			lucro[i] = receita[i] - despesas[i]

			fmt.Println("V: ", z, ", N: ", n[i], ", L: ", lucro[i])
			z++
		}

		fmt.Println("Melhor valor final: ", codML)
		fmt.Println("Lucro: ", maiorLucro)
		fmt.Println("Numero de ingressos: ", maiorNumIngresso)

	} else {
		fmt.Println("INTERVALO INVALIDO")
	}

}
