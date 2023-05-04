package main

import "fmt"

func main() {
	name := "GeraAn"

	switch name {
	case "Gera":
		fmt.Println("Welcome, Gera")
	case "Joko":
		fmt.Println("Welcome, Joko")
	case "Budi":
		fmt.Println("Welcome, Budi")
	default:
		fmt.Println("Welcome home")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}