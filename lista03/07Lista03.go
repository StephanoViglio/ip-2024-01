package main

import (
	"fmt"
)

func freqMap(numSeq []int) map[int]int {

	freq := make(map[int]int)
	for _, num := range numSeq {
		freq[num]++
	}
	return freq
}

func main() {
	var n, m, maiorValor int

	n = 2
	m = 1

	if n > 1 && n <= 10000 {
		for i := 0; m != 0; i++ {

			fmt.Scanln(&n)
			m = n

			numSeq := make([]int, n)

			for k := 0; k < n; k++ {

				fmt.Scan(&numSeq[k])

				for j := 0; j < n; j++ {
					if numSeq[k] > maiorValor {
						maiorValor = numSeq[k]
					}
				}
			}

			freq := freqMap(numSeq)

			for i := 0; i <= maiorValor; i++ {
				for k := 0; k < n; k++ {
					if i > numSeq[k] {
						freq[i]++
					}
				}

			}

			for i := 0; i <= maiorValor; i++ {
				fmt.Println("(", i, ")", freq[i])
			}

			numSeq = nil
			freq = nil

			if n == 0 {
				break
			}
		}

	}
}
