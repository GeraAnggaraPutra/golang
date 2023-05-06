package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.4))
	fmt.Println(math.Floor(1.4))
	fmt.Println(math.Ceil(1.4))
	
	fmt.Println(math.Max(14, 11))
	fmt.Println(math.Min(14, 11))
}