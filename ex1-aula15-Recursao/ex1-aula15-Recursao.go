package main

import "fmt"

func potencia(x, n int) int {

	var pot int

	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		pot = potencia(x, n/2)

		return pot * pot
	} else {
		pot = potencia(x, (n-1)/2)

		return x * pot * pot
	}
}

func main() {

	var x, n int

	fmt.Scanln(&x, &n)

	operacao := potencia(x, n)

	fmt.Println(operacao)

}
