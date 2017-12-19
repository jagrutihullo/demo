package main

import "fmt"

func main() {
	//var i, j int = 0, 0
	var arr = [10]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var s = arr[:4]
	fmt.Println(s, len(s), cap(s))
	s = append(s, 5, 6, 7, 8)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(arr)
	for i := range arr {
		fmt.Scanln(&arr[i])
	}

}
