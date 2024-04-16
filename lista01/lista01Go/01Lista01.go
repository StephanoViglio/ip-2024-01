package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Scanln(&n1, &n2, &n3)
	var x float64 = (n1 + n2 + n3) / 3
	if x >= 6 {
		fmt.Println(x)
		fmt.Println("APROVADO")
	} else {
		fmt.Println(x)
		fmt.Println("REPROVADO")
	}
}
