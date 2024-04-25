package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanln(&n)

	var quadroEscuro int = 2 * n
	var quadroClaro int = n

	var total int = quadroEscuro*32 + quadroClaro*32 - n

	fmt.Println(total)
}
