package main

import "fmt"

func main() {
	x := 15
	a := &x         // & is the pointer in Go (i.e. memory address - will not print out 15, but rather the memory hash)
	fmt.Println(a)  // prints out the memory addy has has of the x
	fmt.Println(*a) // reads through the memory addy to 15
}
