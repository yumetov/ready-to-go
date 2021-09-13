package main

import "fmt"

func main() {
	var age1 int = 45
	fmt.Println(age1 - 28)

	var age2 int64 = 100058
	fmt.Println(age2)
	// fmt.Println(age1 + age2) (mismatched types int and int64)

	fmt.Printf("%T\n", age1)
	fmt.Printf("%T\n", age2)
	fmt.Printf("%T\n", int32(age1))
}