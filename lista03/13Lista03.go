package main

import (
	"fmt"
)

func main() {
	var n, numRep, maiorRep, maiorFreq int

	fmt.Scanln(&n)

	numSeq := make([]int, n)
	freq := make([]int, n)

	if n >= 1 && n <= 100000 {
		for i := 0; i < n; i++ {
			fmt.Scanln(&numSeq[i])
		}

		for i := 0; i < n; i++ {
			if numSeq[i] > 0 && numSeq[i] <= 100 {
				for j := 0; j < n; j++ {
					if numSeq[i] == numSeq[j] {
						freq[i]++
					}

					if freq[i] > maiorFreq {
						maiorFreq = freq[i]
						numRep = numSeq[i]
					}
					if numRep > maiorRep {
						maiorRep = numRep
					}
				}
			}

		}
	}

	fmt.Println(numRep)
	fmt.Println(maiorFreq)
}
