package main

import (
	"fmt"
)

func main() {

	var n, k, x, indice int

	x = 0
	indice = 0

	fmt.Scanln(&n, &k)

	tempoChegada := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&tempoChegada[i])

		if tempoChegada[i] > 0 {
			x++
		}
	}
	if x < k {
		fmt.Println("SIM")
	}
	if x >= k {
		fmt.Println("NÃO")

		for k := n - 1; k >= 0; k-- {
			if tempoChegada[k] <= 0 {
				indice = k + 1
				fmt.Println(indice)
			}
		}
	}

}
