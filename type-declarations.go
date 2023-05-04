package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var gera NoKTP = "768555588098908"
	var marriedStatus Married = false
	fmt.Println(gera)
	fmt.Println(marriedStatus)
}