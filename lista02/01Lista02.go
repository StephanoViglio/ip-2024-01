package main

import (
	"fmt"
)

func main() {

	const qtdProvas = 8
	const qtdListas = 5
	const minPresença = 96

	var matriculas [99]float64
	var trabalhoFinal float64
	var presença float64
	var provas [8]float64
	var somaProvas float64
	var listas [5]float64
	var somaListas float64
	var mediaProvas float64
	var mediaListas float64
	var notaFinal [99]float64

	x := 0
	for i := 0; i < len(matriculas); i++ {

		x++
		fmt.Scan(&matriculas[i])

		if matriculas[i] == -1 {
			break
		}

		for i := 0; i < qtdProvas; i++ {
			fmt.Scan(&provas[i])
			somaProvas += (provas[i])
			mediaProvas = somaProvas / qtdProvas
		}

		for i := 0; i < qtdListas; i++ {
			fmt.Scan(&listas[i])
			somaListas += (listas[i])
			mediaListas = somaListas / qtdListas
		}

		fmt.Scanln(&trabalhoFinal, &presença)

		notaFinal[i] = 0.7*(mediaProvas) + 0.15*(mediaListas) + 0.15*float64(trabalhoFinal)

		fmt.Print("Matricula: ", matriculas[i], " Nota Final: ", notaFinal[i])

		switch {

		case notaFinal[i] >= 6 && presença >= minPresença:
			fmt.Println(" Situação Final: APROVADO")
		case notaFinal[i] >= 6 && presença < minPresença:
			fmt.Println(" Situação Final: REPROVADO POR FREQUENCIA")
		case notaFinal[i] < 6 && presença > minPresença:
			fmt.Println(" Situação Final: REPROVADO POR NOTA")
		case notaFinal[i] < 6 && presença < minPresença:
			fmt.Println(" Situação Final: REPROVADO POR NOTA E POR FREQUENCIA")
		default:
			fmt.Println("VALORES INCORRETOS")
		}
	}

}
