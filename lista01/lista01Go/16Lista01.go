package main

import (
	"fmt"
	"math"
)

func main() {
	var salarioFuncionario float64

	fmt.Scanln(&salarioFuncionario)

	if salarioFuncionario <= 300 {
		var salarioComReajuste float64 = (1.5 * salarioFuncionario)
		fmt.Println("SALARIO COM REAJUSTE = ", math.Round(salarioComReajuste*100)/(100))
	} else {
		var salarioComReajuste float64 = (1.3 * salarioFuncionario)
		fmt.Println("SALARIO COM REAJUSTE = ", math.Round(salarioComReajuste*100)/(100))
	}
}
