package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scanln(&m)

	for n := 1; n <= m; n++ {
		fmt.Printf("\n %d * %d * %d = ", n, n, n)

		min := (n * n) - (n - 1)
		max := (n * n) + (n - 1)

		if n == 1 {
			fmt.Print("1")
		}

		if n >= 2 {
			for i := min; i <= max; i = i + 2 {

				numSeq := i

				fmt.Print(" + ", numSeq)
			}
		}

	}

}
