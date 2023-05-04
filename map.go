package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Gera",
		"address": "Bandung",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Learn Golang"
	book["author"] = "Gera"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(len(book))
}