package main

import (
	"fmt"
)

func main() {

	var numeroHoras int

	fmt.Scanln(&numeroHoras)

	var valorTresHoras int = (numeroHoras) / (3)

	if numeroHoras%3 == 0 {
		var x int = (10 * valorTresHoras)
		fmt.Println("O VALOR A PAGAR E = ", x)
	} else {
		if numeroHoras%3 == 1 {
			var x int = ((10 * valorTresHoras) + 5)
			fmt.Println("O VALOR A PAGAR E = ", x)
		} else {
			if numeroHoras%3 == 2 {
				var x int = ((10 * valorTresHoras) + 10)
				fmt.Println("O VALOR A PAGAR E = ", x)
			}

		}
	}
}
