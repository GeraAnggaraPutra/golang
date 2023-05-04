package main

import "fmt"

func main() {
	var name = "Anggara"

	if name == "Gera" {
		fmt.Println("Hello Gera")
	} else if name == "Joko" { 
		fmt.Println("Hallo Joko")
	} else if name == "Andi" { 
		fmt.Println("Hallo Andi")
	}else {
		fmt.Println("Siapa Kamu ?")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}