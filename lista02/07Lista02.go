package main

import (
	"fmt"
)

func main() {

	var numSeq, somaPar, somaImpar, quantPar, quantImpar int

	for {
		fmt.Scan(&numSeq)

		if numSeq == 0 {
			break
		}

		if numSeq%2 == 0 {
			somaPar += numSeq
			quantPar++
		} else {
			somaImpar += numSeq
			quantImpar++
		}
	}

	mp := somaPar / int(quantPar)
	mi := somaImpar / int(quantImpar)

	fmt.Println("MEDIA PAR = ", mp)
	fmt.Println("MEDIA IMPAR = ", mi)
}
