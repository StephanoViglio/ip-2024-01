package main

import (
	"fmt"
	"math"
)

func main() {

	var numeroJogos int
	fmt.Scanln(&numeroJogos)

	var P float64 //porcentagemPopular
	var G float64 //porcentagemGeral
	var A float64 //porcentagemArquibancada
	var C float64 //porcentagemCadeiras

	T := make([]float64, numeroJogos)
	PP := make([]float64, numeroJogos) //PorcentagemPopular...
	PG := make([]float64, numeroJogos)
	PA := make([]float64, numeroJogos)
	PC := make([]float64, numeroJogos)

	for i := 0; i < numeroJogos; i++ {
		fmt.Scanln(&T[i], &PP[i], &PG[i], &PA[i], &PC[i])
	}

	x := 1
	for i := 0; i < numeroJogos; i++ {
		P = (PP[i]) * T[i]
		G = (PG[i]) * T[i]
		A = (PA[i]) * T[i]
		C = (PC[i]) * T[i]
		var y float64 = (P + 5*G + 10*A + 20*C) / (100)
		fmt.Println("A RENDA DO JOGO Nº ", x, " E = ", math.Round(y*100)/(100))
		x++
	}

}
