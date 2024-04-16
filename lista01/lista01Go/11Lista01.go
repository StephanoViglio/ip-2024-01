package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scanln(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
