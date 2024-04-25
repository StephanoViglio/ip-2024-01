package main

import (
	"fmt"
)

func main() {
	var n, combinacao, numerador int
	var x, y int = 1, 1

	fmt.Scanln(&n)
	f := (n - 2)

	for i := n; i > 1; i-- {
		x = x * i
	}

	for i := f; i > 1; i-- {
		y = y * i
	}

	numerador = x / 2
	combinacao = numerador / y //combinação = k;

	if n > 1 {

		timeA := []int{1, 2, 3}
		timeB := []int{3, 1, 2}

		w := 1
		z := 0
		for i := 0; i < combinacao; i++ {

			fmt.Println("Final ", w, ":", "Time", timeA[i], "X", "Time", timeB[z])
			w++
			z++
		}

	}

}
