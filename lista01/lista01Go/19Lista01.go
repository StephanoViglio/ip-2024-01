package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var y int = 1

	fmt.Scanln(&n)

	var somaFracoes float64
	var x float64 = 1

	if n > 1 {
		for y <= n {
			somaFracoes = somaFracoes + 1/x
			x++
			y++
		}
		fmt.Println(math.Round(somaFracoes*1000000) / (1000000))

	} else {
		fmt.Println("NUMERO INVALIDO!")
	}
}
