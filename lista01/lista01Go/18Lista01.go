package main

import (
	"fmt"
)

func main() {
	var valorInicial, razaoPA, numeroElementos int

	fmt.Scanln(&valorInicial, &razaoPA, &numeroElementos)

	var somaPA int = (2*valorInicial + razaoPA*(numeroElementos-1)) * (numeroElementos) / (2)

	fmt.Println(somaPA)
}
