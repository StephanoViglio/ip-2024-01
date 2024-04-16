package main

import (
	"fmt"
	"math"
)

func main() {
	var valorNota float64

	fmt.Scanln(&valorNota, math.Round(valorNota*10)/10)

	if valorNota >= 9 && valorNota <= 10 {
		fmt.Println("NOTA = ", math.Round(valorNota*10)/10, "CONCEITO = A")
	}
	if valorNota < 9 && valorNota >= 7.5 {
		fmt.Println("NOTA = ", math.Round(valorNota*10)/10, "CONCEITO = B")
	}
	if valorNota < 7.5 && valorNota >= 6.0 {
		fmt.Println("NOTA = ", math.Round(valorNota*10)/10, "CONCEITO = C")
	}
	if valorNota < 6.0 && valorNota >= 0 {
		fmt.Println("NOTA = ", math.Round(valorNota*10)/10, "CONCEITO = D")
	}
}
