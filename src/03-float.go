package main

import "fmt"

func main() {
	var pi float64 = 3.142
	sqrt2 := 1.414
	var e float32 = 2.718

	fmt.Println(pi, sqrt2, e)
	fmt.Printf("%T %T %T %T\n", pi, sqrt2, e, float64(e))

	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := 0.0 / zero
	fmt.Println(zero, pinf, ninf, nan)
	fmt.Printf("%T %T %T %T\n", zero, pinf, ninf, nan)
}