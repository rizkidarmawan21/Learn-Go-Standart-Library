//  Package strconv adalah package yang berisi fungsi-fungsi untuk mengkonversi string ke tipe data lain

package main

import (
	"strconv"
)

func main(){
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Boolean:", boolean)
	}


	intValue, err := strconv.Atoi("123")
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Integer:", intValue)
	}

	binaryValue := strconv.FormatInt(123, 2)
	println("Binary:", binaryValue)
}