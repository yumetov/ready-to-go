package main

import "fmt"

func main() {
	var array1 [3]int = [3]int{1, 2}
	fmt.Println(array1)
	fmt.Printf("%T\n", array1)

	array2 := [...]string{"", "H", "He", "Li"}
	fmt.Println(array2)
	fmt.Printf("%T\n", array2)
	fmt.Println(len(array2))

	array3 := [...]int{1, 2, 3}
	array4 := [...]int{4, 5}

	// type error
	// array3 = array4
	fmt.Println(array3)

	// type error
	// array4 = array3
	fmt.Println(array4)
}
