package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var maxValor float64

	fmt.Scanln(&n)

	pontosXYZ := make([][3]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&pontosXYZ[i][0], &pontosXYZ[i][1], &pontosXYZ[i][2])
	}

	vetorAB := make([][3]float64, n-1)

	for i := 0; i < n-1; i++ {
		vetorAB[i][0] = pontosXYZ[i+1][0] - pontosXYZ[i][0]
		vetorAB[i][1] = pontosXYZ[i+1][1] - pontosXYZ[i][1]
		vetorAB[i][2] = pontosXYZ[i+1][2] - pontosXYZ[i][2]
	}

	for _, valor := range vetorAB {
		maxValor = 0
		for _, coord := range valor {
			if math.Abs(coord) > math.Abs(maxValor) {
				maxValor = coord
			}
		}

		fmt.Printf("%.2f \n", math.Abs(maxValor))
	}

}
