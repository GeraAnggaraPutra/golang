package main

import "fmt"

func main() {
	const firstName string = "Gera"
	const lastName = "Anggara"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		nama = "Gera"
		kota = "Bandung"
	)

	fmt.Println(nama)
	fmt.Println(kota)
}