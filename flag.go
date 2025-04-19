// Flag merupakan struct yang digunakan untuk menyimpan informasi tentang flag

package main

import (
	"flag"
	"fmt"
)


func main(){
	name := flag.String("name", "default", "your name")
	password := flag.String("password", "default", "your password")

	flag.Parse()

	fmt.Println("Name:", *name)
	fmt.Println("Password:", *password)

	// go run flag.go -name=coba -password=123
}