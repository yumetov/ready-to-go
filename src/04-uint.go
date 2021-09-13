package main

import "fmt"

func main() {
	var money uint = 200
	var juice uint = 120
	fmt.Println(money - juice)

	var nmoney uint32
	nmoney = uint32(juice) - uint32(money)
	fmt.Println(nmoney) // 4294967216
}