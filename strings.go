package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Gera Anggara", "Gera"))
	fmt.Println(strings.Contains("Gera Anggara", "Anto"))

	fmt.Println(strings.Split("Gera Anggara Putra", " "))

	fmt.Println(strings.ToLower("Gera Anggara Putra"))
	fmt.Println(strings.ToUpper("Gera Anggara Putra"))
	fmt.Println(strings.ToTitle("Gera Anggara Putra"))

	fmt.Println(strings.Trim("      Gera Anggara     ", " "))
	fmt.Println(strings.ReplaceAll("Gera Joko Gera", "Gera", "Anto"))
}