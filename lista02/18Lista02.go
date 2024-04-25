package main

import (
	"fmt"
)

func main() {

	var x, y, z int

	fmt.Scanln(&x, &y, &z)

	fmt.Println(x, y, z, ":2")

	x2, y2, z2 := x/2, y/2, z/2

	switch {
	case x%2 == 0 && y%2 == 0 && z%2 == 0:
		fmt.Print(x2, y2, z2)
	case x%2 == 0 && y%2 == 0 && z%2 != 0:
		fmt.Print(x2, y2, z)
	case x%2 == 0 && y%2 != 0 && z%2 != 0:
		fmt.Print(x2, y, z)
	case x%2 == 0 && y%2 != 0 && z%2 == 0:
		fmt.Print(x2, y, z2)
	case x%2 != 0 && y%2 == 0 && z%2 == 0:
		fmt.Print(x, y2, z2)
	default:
		fmt.Println(x, y, z, ":3")

	}
}
