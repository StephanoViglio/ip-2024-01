package main

import (
	"fmt"
)

func main() {

	var m, n int

	fmt.Scanln(&m, &n) //m linhas e n colunas

	var colunas [999]int

	fmt.Println("(2, 1)")

	for j := 1; j < n; j++ {
		colunas[j] += j
		fmt.Printf("(%d, %d)", m, colunas[j])
	}

}
