package main

import (
	"fmt"
)

func main() {

	var m = map[int]int{
		1: 3,
		2: 4}
	//m = make(map[int]int)
	//m[1] = 1
	//fmt.Println(m[1])
	var key int
	fmt.Scanln(&key)
	ele, ok := m[key]
	fmt.Println(ele, ok)
}
