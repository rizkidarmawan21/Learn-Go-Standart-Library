package main

import (
	"fmt"
	"regexp"
)

func main(){
	regexp := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regexp.MatchString("Eko"))
	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.FindAllString("eko ebo eKu",10))
}