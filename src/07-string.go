package main

import "fmt"

func main() {
	var shout string = "Ready to \"Go\""
	fmt.Println(shout)
	fmt.Printf("%T\n", shout)

	var sushi string = `Tsuna: 990 yen
Mackerel: 660 yen
Squid: 880 yen`
	fmt.Println(sushi)

	var introduction string = "Her name is Chihaya."
	fmt.Println(introduction, introduction[0])
	fmt.Printf("%T\n", introduction[0])
}