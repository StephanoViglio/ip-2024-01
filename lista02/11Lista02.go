package main

import (
	"fmt"
)

func main() {
	var n, e float64

	fmt.Scanln(&n)
	fmt.Scanln(&e)

	rk := make([]float64, 100)

	for i := 1; i < 100; i++ {

		rk[0] = 1
		rk[i] = (rk[i-1] + n/(rk[i-1])) / (2)

		dif := (n - rk[i]*rk[i]) * (-1)

		fmt.Println("r: ", rk[i], ", erro: ", dif*(1000000000)/(1000000000))

		if dif < e {
			break
		}

	}

}
