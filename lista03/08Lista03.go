package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scanln(&n) //quantidade de números a serem escritos

	numSeq := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&numSeq[i])

	}

	for j := 0; j < n; j++ {
		fmt.Println(strconv.FormatInt(numSeq[j], 2))
	}

}
