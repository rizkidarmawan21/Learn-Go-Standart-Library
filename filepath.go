package main

import (
	"fmt"
	"path/filepath"
)

func main(){
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Dir("hello/guys/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go")) // artinya ambil dari depan
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello","world", "main.go"))
}