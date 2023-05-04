package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Gera"
	names[1] = "Anggara"
	names[2] = "Putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values)) 
	fmt.Println(len(names)) 

	var test [10]string
	fmt.Println(len(test))
}