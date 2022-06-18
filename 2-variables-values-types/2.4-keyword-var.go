package main

import (
	"fmt"
)

func main() {

	// Only code-block can see the variable below
	y := 10
	qualquercoisa(y)
}

func qualquercoisa(x int) {
	fmt.Println(x)
	// fmt.Println(y) It doesn't work because y is coded-block
}

// Para declarar variaveis com package level scope é necessário
// atribuir o var e fora de um code block.
