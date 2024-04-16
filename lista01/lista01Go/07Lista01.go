package main

import (
	"fmt"
	"math"
)

func main() {
	var valorFahrenheit, quantChuvaPolegadas float64
	fmt.Scanln(&valorFahrenheit)
	fmt.Scanln(&quantChuvaPolegadas)

	var valorCelsius float64 = (5) * (valorFahrenheit - 32) / (9)
	var valorChuvaMili float64 = (25.4) * (quantChuvaPolegadas)

	fmt.Println("O VALOR EM CELSIUS = ", math.Round(valorCelsius*100)/(100))
	fmt.Println("A QUANTIDADE DE CHUVA E = ", math.Round(valorChuvaMili*100)/(100))
}
