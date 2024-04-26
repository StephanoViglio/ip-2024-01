package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	fmt.Print(n, " = ")

	soma := 0
	for i := 1; i < n; i++ {

		if n%i == 0 {
			soma = soma + i
			fmt.Printf("+ %d", i)
		}
	}
	if soma == n {
		fmt.Printf("= %d (NUMERO PERFEITO)", soma)
	} else {
		fmt.Printf("= %d (NUMERO NAO E PERFEITO)", soma)
	}
}
