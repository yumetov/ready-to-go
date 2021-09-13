package main

import "fmt"

func main() {
	var c1 complex64 = -4 + 3i
	var c2 complex64 = 1 + 2i
	fmt.Println(c1, c2, c1+c2)
}