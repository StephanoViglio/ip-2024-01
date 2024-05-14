package main

import (
	"fmt"
	"math"
)

func main() {
	var t int

	fmt.Scanln(&t)

	vetor := make([]int, 11)

	for i := 0; i < t; i++ {

		for k := 0; k < 11; k++ {
			fmt.Scan(&vetor[k])
		}

		somaIda := vetor[0]*1 + vetor[1]*2 + vetor[2]*3 + vetor[3]*4 + vetor[4]*5 + vetor[5]*6 + vetor[6]*7 + vetor[7]*8 + vetor[8]*9

		resto1 := math.Remainder(float64(somaIda), 11)
		b1 := 11 + resto1

		somaVolta := vetor[0]*9 + vetor[1]*8 + vetor[2]*7 + vetor[3]*6 + vetor[4]*5 + vetor[5]*4 + vetor[6]*3 + vetor[7]*2 + vetor[8]*1

		resto2 := math.Remainder(float64(somaVolta), 11)
		b2 := 11 + resto2

		if vetor[9] == int(b1) && vetor[10] == int(b2) {
			fmt.Println("CPF VALIDO")
		} else {
			fmt.Println("CPF INVALIDO")
		}
	}

}
