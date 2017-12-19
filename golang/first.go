package main

import (
	"fmt"
)

var s string = "global"

func add(i, j int) (x, y int) {
	x = i + j
	y = i - j
	return
}

func main() {
	k, l := add(5, 7)
	const i = 56.53435
	const j = i
	fmt.Printf("%g %T\n", j, j)
	//	fmt.Printf("%d %d\n", k, l)
	//	fmt.Println("string " + s)
}
