package main

import "fmt"

func main() {
	b1 := []byte{71, 111, 33}
	fmt.Println(b1)
	fmt.Println(string(b1))

	b2 := []byte("Go!")
	fmt.Println(b2)
	fmt.Println(string(b2))
}