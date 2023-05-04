package main

import "fmt"

func main() {
	name := "Gera"
	counter := 0

	increment := func() {
		name := "Anto"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}