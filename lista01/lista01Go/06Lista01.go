package main

import (
	"fmt"
	"math"
)

func main() {

	var numeroTemperaturas int
	fmt.Scanln(&numeroTemperaturas)

	x := make([]float64, numeroTemperaturas) //valorFahrenheit

	for i := 0; i < numeroTemperaturas; i++ {
		fmt.Scanln(&x[i])
	}

	for i := 0; i < numeroTemperaturas; i++ {
		y := (5) * (x[i] - 32) / (9)
		fmt.Println(x[i], " FAHRENHEIT EQUIVALE A ", math.Round(y*100)/(100), " CELSIUS")
	}

}
