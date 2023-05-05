package helper

import "fmt"

// jika namanya diawali dengan huruf besar maka artinya bisa di akses dari luar, tapi jika kecil maka tidak bisa

var App = "Belajar Golang" // bisa diakses dari luar
var version = "1.0.0" // tidak bisa diakses dari luar

func sayGoodbye(name string) { // tidak bisa diakses dari luar
	fmt.Println("Hello " + name)
}

func SayHello(name string) { // bisa diakses dari luar
	fmt.Println("Hello " + name)
}