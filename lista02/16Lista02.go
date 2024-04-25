package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for h := 1; h <= n; h++ {
		for c1 := 1; c1 < h; c1++ {
			for c2 := c1; c2 < h; c2++ {
				if h*h == c1*c1+c2*c2 {
					fmt.Println("hipotenusa = ", h, " catetos: ", c1, " e ", c2)
				}
			}
		}
	}
}
