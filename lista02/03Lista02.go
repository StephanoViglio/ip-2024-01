package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Scanln(&n)

	var x float64 = 1
	for i := n; i > 1; i-- {
		x = x * i
	}
	fmt.Println(n, "! = ", x)
}
