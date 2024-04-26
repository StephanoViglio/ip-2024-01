package main

import (
	"fmt"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {

	var x, y, z int
	fmt.Scanln(&x, &y, &z)

	switch {
	case x%2 == 0 && y%2 == 0 && z%2 == 0:
		fmt.Print(x/2, y/2, z/2)
	case x%2 == 0 && y%2 == 0 && z%2 != 0:
		fmt.Print(x/2, y/2, z)
	case x%2 == 0 && y%2 != 0 && z%2 != 0:
		fmt.Print(x/2, y, z)
	case x%2 == 0 && y%2 != 0 && z%2 == 0:
		fmt.Print(x/2, y, z/2)
	case x%2 != 0 && y%2 == 0 && z%2 == 0:
		fmt.Print(x, y/2, z/2)
	default:
		fmt.Println(x, y, z, ":3")

	}

	fmt.Println("\n MMC: ", LCM(x, y, z))
}
