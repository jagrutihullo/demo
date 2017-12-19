package main

import "fmt"

func main() {
	var i int = 0
	for i < 20 {
		fmt.Print(i)
		i++
	}
	fmt.Printf("%d\n", i)
}
