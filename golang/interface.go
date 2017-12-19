package main

import (
	"fmt"
)

type mytype rune

type myinterface interface {
	abs() float32
}

func main() {
	var m myinterface
	f := mytype(5)

	m = f
	fmt.Println(m.abs())
	fmt.Printf("%v %T\n", m, m)
}

func (a mytype) abs() float32 {
	return float32(a)
}
