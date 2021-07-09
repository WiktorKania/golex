package main

import (
	"ckript/token"
	"fmt"
)

func main() {
	var t = token.New(token.ALLOC, "Value", 0, "")
	fmt.Println("Hello", t)
}
