package main

import (
	"fmt"
	"math"
)

func main() {

	var A, B, C float64

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	var x float64 = (B*B - 4*A*C)

	fmt.Println("O VALOR DE DELTA E = ", math.Round(x*100)/(100))
}
