package main

import (
	"fmt"
	"learn-golang/database" 
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}