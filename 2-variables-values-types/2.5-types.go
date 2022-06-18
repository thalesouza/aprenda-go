// Types are the base that what Go is.

// Types in go are static. If you need another variable, declare
// another.

package main

import (
	"fmt"
)

var x int = 10

func main() {

	x = 20.5
	fmt.Println(x)
}
