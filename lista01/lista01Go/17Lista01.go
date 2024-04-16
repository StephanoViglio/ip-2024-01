package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scanln(&x, &y)

	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Printf("%v ", x)
			x = x + 2
		}

	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}
