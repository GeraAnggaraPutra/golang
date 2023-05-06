package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")
	
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	var regex2 *regexp.Regexp = regexp.MustCompile("g([c-g][a-z])a")
	fmt.Println(regex2.MatchString("gera"))
	fmt.Println(regex2.MatchString("jera"))

	search := regex.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)
}