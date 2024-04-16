package main

import (
	"fmt"
	"math"
)

func main() {
	var salarioMin, quantKwCons float64

	fmt.Scanln(&salarioMin)
	fmt.Scanln(&quantKwCons)

	var valorKw float64 = (7) * (salarioMin) / (1000)

	var consumoResid float64 = (valorKw) * (quantKwCons)
	var novoValorCons float64 = (consumoResid) * (9) / (10)

	fmt.Println("Custo por kW: R$", math.Round(valorKw*100)/100)
	fmt.Println("Custo do consumo: R$", math.Round(consumoResid*100)/100)
	fmt.Println("Custo com desconto: R$", math.Round(novoValorCons*100)/100)
}
