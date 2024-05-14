package main

import (
	"fmt"
)

func main() {
	var n, k, x int

	fmt.Scanln(&n)

	vetorV := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&vetorV[i])
	}
	fmt.Scanln(&k)

	for _, maiorValor := range vetorV {
		if maiorValor >= k {
			x++
		}
	}
	fmt.Println(x)
}
