package main

import (
	"fmt"
)

func main() {
	var n int
	var vetorV [1000]int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vetorV[i])
		fmt.Print(vetorV[i], " ")
	}

}
