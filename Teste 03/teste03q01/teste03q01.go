package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var frase string
	for frase != "#" {
		reader := bufio.NewReader(os.Stdin)
		frase, _ := reader.ReadString('\n')
		frase = strings.TrimSpace(frase)
		if frase == "#" {
			break
		}
		split := strings.Split(frase, "")
		for i := len(split) - 1; i > 0; i-- {
			fmt.Print(split[i])
		}
	}
}
