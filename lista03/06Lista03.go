package main

import (
	"fmt"
)

func main() {
	var n int
	var vetorV [5000]int

	fmt.Scanln(&n)

	soma := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vetorV[i])
		soma = soma + vetorV[i]
	}
	fmt.Println(soma)
}
