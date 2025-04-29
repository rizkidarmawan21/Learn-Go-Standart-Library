package main

import (
	"fmt"
	"slices"
)

func main(){
	names := []string{"Alice", "Bob", "Charlie"}
	values := []int{1, 2, 3}


	fmt.Println(slices.Min(names)) // Output: Alice
	fmt.Println(slices.Min(values)) // Output: 1
	fmt.Println(slices.Max(names)) // Output: Charlie
	fmt.Println(slices.Max(values)) // Output: 3
	fmt.Println(slices.Contains(names, "Bob")) // Output: true
	fmt.Println(slices.Contains(names, "Eve")) // Output: false
	fmt.Println(slices.Index(names, "Bob")) // Output: 1
}