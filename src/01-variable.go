package main

import "fmt"

func main(){
	var name string = "Aoba"
	fmt.Println(name)

	var age int = 18
	fmt.Println(age)

	var adult, leader bool = false, true
	fmt.Println(adult, leader)

	var (
		team string = "character"
		wyear = 2
	)
	fmt.Println(team, wyear)

	var height int
	fmt.Println(height)
	height = 149
	fmt.Println(height)

	birthday := "0202"
	fmt.Println(birthday)
	// birth = 202 (error)
}
