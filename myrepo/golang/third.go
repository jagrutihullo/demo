package main

import (
	"fmt"
)

type demo struct {
	x float32
	y int64
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	var p *demo
	// var err error
	var d demo
	// //fmt.Print("Enter x: ")
	// d.x = "83747"
	// d.y, err = strconv.ParseInt(d.x, 10, 32)
	// fmt.Print(d, "\n", err)

	p = &d
	fmt.Scanln(&d.x)
	fmt.Scanln(&d.y)
	fmt.Print(p, "\n")
}
