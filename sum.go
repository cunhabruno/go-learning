package main

import "fmt"

func sum(x, y int) int {
	fmt.Println("My go lang sum function :)")

	result := x + y
	fmt.Println("x+y =", result)
	return result
}
