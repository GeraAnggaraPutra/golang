package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 20
	var c = a * b
	fmt.Println(c)

	// augmented assignments
	var i = 10
	i += 40 // i = i + 40
	fmt.Println(i)

	// unary operator
	i++ // i = 1 + 1
	fmt.Println(i)
	
}