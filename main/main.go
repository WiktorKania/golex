package main

import (
	"fmt"
	"golex/token"
)

func main() {
	var t = token.New(token.ALLOC, "Value", 0, "")
	fmt.Println("Hello", t)
}
