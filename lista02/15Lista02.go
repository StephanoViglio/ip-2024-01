package main

import (
	"fmt"
	"strconv"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func compararNum(x, y int) int {

	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)

	reversedX := reverseString(strX)
	reversedY := reverseString(strY)

	reversedIntX, _ := strconv.Atoi(reversedX)
	reversedIntY, _ := strconv.Atoi(reversedY)

	if reversedIntX > reversedIntY {
		return reversedIntX
	} else {
		return reversedIntY
	}
}

func main() {

	var t int
	fmt.Scanln(&t)

	num1 := make([]int, t)
	num2 := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&num1[i], &num2[i])
	}
	for i := 0; i < t; i++ {

		resultComp := compararNum(num1[i], num2[i])

		fmt.Println(resultComp)
	}
}
