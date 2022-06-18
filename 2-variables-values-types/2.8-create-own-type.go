package main

import "fmt"

// when creating own type we have more options than just using the built-in types
type hotdog int

var b hotdog

func main() {
	fmt.Printf("%T", b)
}

// the subjacent type doesn't create the relationship between the two types
