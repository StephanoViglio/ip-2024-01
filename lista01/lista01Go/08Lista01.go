package main

import (
	"fmt"
	"math"
)

func main() {
	var raioLata, alturaLata float64

	fmt.Scanln(&raioLata)
	fmt.Scanln(&alturaLata)

	var π float64 = 3.14159
	var Ac float64 = (π) * (raioLata) * (raioLata)
	var Al float64 = (2 * π) * (raioLata) * (alturaLata)
	var At float64 = ((2 * Ac) + Al)
	var x float64 = (100 * At)

	fmt.Println("O VALOR DO CUSTO E = ", math.Round(x*100)/(100))
}
