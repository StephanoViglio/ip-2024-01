package main

import (
	"fmt"
)

func somaArray(array []float64, i int) float64 {

	if i >= len(array) {
		return 0
	}

	return array[i] + somaArray(array, i+1)

}

func main() {

	var n int
	var array [999]float64

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&array[i])
	}

	somaTotal := somaArray(array[:], 0)

	fmt.Println(somaTotal)

}
