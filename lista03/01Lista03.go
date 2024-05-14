package main

import (
	"fmt"
)

func main() {
	var n, m, x int

	fmt.Scanln(&n)

	vetorV := make([]int, n)

	if n >= 1 && n <= 100000 {
		for i := 0; i < n; i++ {
			fmt.Scanln(&vetorV[i])
		}
	}

	fmt.Scanln(&m)
	numeros := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanln(&numeros[i])
	}
	for i := 0; i < m; i++ {
		for k := 0; k < n; k++ {
			if numeros[i] == vetorV[k] {
				fmt.Println("ACHEI")
				x++
			}
		}
		if x == 0 {
			fmt.Println("NAO ACHEI")
		}
		x = 0
	}
}
