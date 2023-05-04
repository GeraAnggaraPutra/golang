package main

import "fmt"

func getFullName() (string, string) {
	return "Gera", "Anggara"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
}