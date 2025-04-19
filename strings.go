// Strings adalah package yang berisi fungsi-fungsi untuk memanipulasi string
// example: strings.Contains, strings.Split, strings.Join, dll

package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Hello, World!", "World")) // true
	fmt.Println(strings.Contains("Hello, World!", "world")) // false

	fmt.Println(strings.Split("Hello, World!", " ")) // [Hello, World!]
	fmt.Println(strings.ToLower("Hello, World!")) // hello, world!
	fmt.Println(strings.ToUpper("Hello, World!")) // HELLO, WORLD!
	fmt.Println(strings.TrimSpace(" Hello, World! ")) // Hello, World!
	fmt.Println(strings.Trim("Hello, World!", "Helo")) // , World!
	fmt.Println(strings.ReplaceAll("Hello, World!", "World", "Golang")) // Hello, Golang!
}

