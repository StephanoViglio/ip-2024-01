package main

import (
	"fmt"
)

func main() {
	var tempoHoras, tempoMinutos, tempoSegundos int

	fmt.Scanln(&tempoHoras)
	fmt.Scanln(&tempoMinutos)
	fmt.Scanln(&tempoSegundos)

	var tempoTotalSegundos int = (3600*tempoHoras + 60*tempoMinutos + tempoSegundos)

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d", tempoTotalSegundos)
}
