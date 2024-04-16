package main

import (
	"fmt"
	"math"
)

func main() {
	var alturaPiramide, arestaHexagono float64

	fmt.Scanln(&alturaPiramide)
	fmt.Scanln(&arestaHexagono)

	var sqrt3 float64 = 1.73
	var volumePiramide float64 = (3 * (arestaHexagono * arestaHexagono) * alturaPiramide * sqrt3)

	fmt.Println("O VOLUME DA PIRAMIDE E = ", math.Round(volumePiramide*100)/(100))
}
