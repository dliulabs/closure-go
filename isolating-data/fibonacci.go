package main

import (
	"fmt"
)

func makeFibGen() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		// must defer calculations to return the un-modified f1
		defer func() { 
			f2, f1 = (f1 + f2), f2 // should set f2 first
		}()
		return f1
	}
}

func main() {
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Print(gen(), ", ")
	}
}
