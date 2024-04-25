package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a < b {

		y := math.Log(b / a)
		w := math.Log(1.03 / 1.015)
		x := (y) / (w)

		if x > 0 {
			fmt.Println("ANOS = ", math.Ceil(x))
		}

	} else {
		fmt.Println("Valores de população incorretos!")
	}

}
