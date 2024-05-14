package main

import (
	"fmt"
)

func removerDuplicado(numSeq []int) []int {
	map_elementos := map[int]bool{}
	result := []int{}

	for i := range numSeq {
		if !map_elementos[numSeq[i]] {
			map_elementos[numSeq[i]] = true
			result = append(result, numSeq[i])
		}
	}
	return result
}

//Com a criação do mapa, irá se analisar cada elemento do array, e a parte "!map_elementos[numSeq[i]]" (que equivale a "!= true")...
//...checa se o elemento já está no mapa. Se ele não está presente (ao verificar que não é true), é porque...
//...ele está encontrando o elemento pela primeira vez, marcando-o como true e indicando que ele foi verificado...
//...e colocado (adicionado/append) no "result", que será retornado no final. Isso ocorre até todos os elementos...
//...do array forem verificados e os únicos colocados no "result".

func main() {
	var n int

	fmt.Scanln(&n)

	numSeq := make([]int, n)

	if n >= 1 && n <= 1000 {
		for i := 0; i < n; i++ {
			fmt.Scanln(&numSeq[i])
		}

		result := removerDuplicado(numSeq)

		fmt.Println(result)
	}
}
