package main

import (
	"fmt"
)

func generator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	numGenerator := generator() // returns a 'func() int'
	for i := 0; i < 5; i++ {
		fmt.Println(numGenerator(), "\t")
	}

}