package main

import (
	"fmt"
)

func main() {
	var n, maiorV, menorW int

	menorW = 1000

	fmt.Scanln(&n)

	vetorV := make([]int, n)

	if n > 1 && n <= 1000 {

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &vetorV[i])
			fmt.Print(vetorV[i], " ")
		}

		for i := 0; i < n; i++ {
			if vetorV[i] > maiorV {
				maiorV = vetorV[i]
			}

			if vetorV[i] < menorW {
				menorW = vetorV[i]
			}
		}
	}
	fmt.Println()

	for i := n - 1; i >= 0; i-- {
		fmt.Print(vetorV[i], " ")
	}
	fmt.Println()
	fmt.Println(maiorV)
	fmt.Println(menorW)
}
