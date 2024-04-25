package main

import (
	"fmt"
	"math"
)

func checarPrimo(n int) {

	sqRoot := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqRoot; i++ {
		if n%i == 0 {
			fmt.Println("NAO PRIMO")
			return
		}
	}
	fmt.Println("PRIMO")
}

func main() {
	var n int
	fmt.Scanln(&n)

	checarPrimo(n)
}
