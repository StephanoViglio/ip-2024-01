package main

import (
	"fmt"
)

func main() {

	var m, n int

	fmt.Scanln(&m, &n) //m linhas e n colunas

	linhas := make([]int, m+1)
	colunas := make([]int, n+1)

	fmt.Println("(2, 1)")

	for i := 1; i < m; i++ {

		linhas[i] += i

		colunas[i] += i
		fmt.Printf("(%d, %d)", m, colunas[i])

	}

}
