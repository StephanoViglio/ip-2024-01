package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanln(&n)

	if n > 5 && n < 2000 && n%2 == 0 {
		for x := 2; x <= n; x = x + 2 {
			fmt.Println(x, "^2 = ", x*x)
		}
	}
	if n > 5 && n < 2000 && n%2 != 0 {
		for x := 2; x < n; x = x + 2 {
			fmt.Println(x, "^2 = ", x*x)
		}
	}
}
