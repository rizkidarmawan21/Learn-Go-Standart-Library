package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Rizki Darmawan"

	encoding := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Print("Encoding: ", encoding, "\n")

	decoded, err := base64.StdEncoding.DecodeString(encoding)
	if err != nil {
		fmt.Println("Error decoding:", err)
	}
	fmt.Print("Decoded: ", string(decoded), "\n")
}