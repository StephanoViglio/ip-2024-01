package main

import (
	"fmt"
)

func main() {

	matricula := make([]float64, 100)
	horas := make([]float64, 100)
	valorHora := make([]float64, 100)

	x := 0
	for i := 0; i < 100; i++ {
		fmt.Scanln(&matricula[i], &horas[i], &valorHora[i])
		x++

		if matricula[i] == 0 && horas[i] == 0 && valorHora[i] == 0 {
			break
		}
	}

	for i := 0; i < x-1; i++ {

		salario := horas[i] * valorHora[i]
		fmt.Println(matricula[i], salario)
	}

}
