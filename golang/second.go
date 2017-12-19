package main

import "fmt"

func main() {
	var i int = 0
	for {
		fmt.Print(i)
		i++
	}
	fmt.Printf("%d\n", i)
}
