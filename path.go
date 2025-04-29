// digunakan untuk memanipulasi path di url/di file sistem. Pada dasarnya menggunakan slash.

package main

import (
	"fmt"
	"path"
)

func main(){
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Dir("hello/guys/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello","world", "main.go"))
}