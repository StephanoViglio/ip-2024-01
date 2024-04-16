package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	var x float64 = (a*d - b*c)

	fmt.Println("O VALOR DO DETERMINANTE E = ", math.Round(x*100)/(100))

}
