package main

// Package fmt (format) implements formatted I/O with functions analogous to C's printf and scanf
// if we have vars or imports we dont use, it will throw and error
import (
	"fmt"
) // need to use double quotes around imports (all double quotes, unlike python which is single)

// //Print to console
// func main() {
// 	fmt.Println("Welcome to GO") // Very Java based feel
// }

// func foo() {
// 	fmt.Println("The square root of 4 is", math.Sqrt(4)) // Caps matters in GO, if the function is caps, it will be exported by go, it not, it will be an internal (i.e. public, private)
// }

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) { // To return an array, we must specify the return type in the brackets like this (i.e. a and b are both strings)
	return a, b
}

// Same with Java, main will always run
func main() {
	num1, num2 := 5.6, 9.5 // Note: you dont have to specify type inside a function, go will figure out when you compile, but then the type will not be changeable (with colon) (static) - but not good, try to be explicit
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"

	var a int = 62
	var b float64 = float64(a) // convert types

	// inference type
	x := a // x will be type int

	fmt.Println(multiple(w1, w2))
}
