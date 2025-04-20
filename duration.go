package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Duration(1 * time.Second)
	var duration2 time.Duration = 10 * time.Millisecond

	fmt.Println("Duration 1:", duration1)
	fmt.Println("Duration 2:", duration2)
}