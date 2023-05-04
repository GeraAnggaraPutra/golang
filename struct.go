package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var gera Customer
	gera.Name = "Gera"
	gera.Address = "Indonesia"
	gera.Age = 17

	gera.sayHi("Joko")
	gera.sayHuuu()

	// fmt.Println(gera.Name)
	// fmt.Println(gera.Address)
	// fmt.Println(gera.Age)
	
	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Cirebon",
	// 	Age:     35,
	// }
	// fmt.Println(joko)
	
	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)
}