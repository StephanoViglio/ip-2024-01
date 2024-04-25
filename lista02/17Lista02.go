package main

import (
	"fmt"
)

func main() {

	var codigo, preçoCompra, preçoVenda, numVendas, porcentagemLucro [9]float64
	var lucroMenorQue10, lucroEntre10e20, lucroMaiorQue20, maiorLucro, maisVendida, codML, valorTotal, valorTotalCompras float64

	for i := 0; i < 5; i++ {
		fmt.Scan(&codigo[i], &preçoCompra[i], &preçoVenda[i], &numVendas[i])

		porcentagemLucro[i] = preçoVenda[i] / preçoCompra[i]
		valorTotal = preçoCompra[i] * numVendas[i]
		valorTotalCompras += valorTotal

		if codigo[i] == -1 && preçoCompra[i] == -1 && preçoVenda[i] == -1 && numVendas[i] == -1 {
			break
		}
		if porcentagemLucro[i] < 1.1 {
			lucroMenorQue10++
		} else if porcentagemLucro[i] >= 1.1 && porcentagemLucro[i] <= 1.2 {
			lucroEntre10e20++
		} else {
			lucroMaiorQue20++
		}

		if porcentagemLucro[i] > maiorLucro {
			maiorLucro = porcentagemLucro[i]
			codML = codigo[i]
		}
		if numVendas[i] > maisVendida {
			maisVendida = codigo[i]
		}
	}

	fmt.Println("Quantidade de mercadorias que geraram lucro menor que 10%: ", lucroMenorQue10)
	fmt.Println("Quantidade de mercadorias que geraram lucro maior ou igual a 10%, e menor ou igual a 20%: ", lucroEntre10e20)
	fmt.Println("Quantidade de mercadorias que geraram lucro maior do que 20%: ", lucroMaiorQue20)
	fmt.Println("Codigo da mercadoria que gerou maior lucro: ", codML)
	fmt.Println("Codigo da mercadoria mais vendida: ", maisVendida)
	fmt.Println("Valor total de compras: ", valorTotalCompras)

}
