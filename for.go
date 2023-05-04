package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++
	}

	for c := 1; c <= 5; c++ {
		fmt.Println("Perulangan -", c)
	}

	slice := []string{"Gera", "anggara", "putra"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i,"=",value)
	}

	person := make(map[string]string)
	person["name"] = "Gera"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key," = ",value)
	}
}