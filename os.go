// Package OS adalah package yang berisi fungsi-fungsi untuk berinteraksi dengan sistem operasi
// Seperti membuka file, read file, write file, dll

package main

import "os"

// os.Argv adalah array yang berisi argumen yang diberikan saat menjalankan program
// contoh ketika menjalankan program go run main.go arg1 arg2
// maka os.Args[0] = "main.go", os.Args[1] = "arg1", os.Args[2] = "arg2"
// dibelakang dia dicompile menjadi binary makanya array 1 itu adalah lokasi file hasil compile

func main() {
	args := os.Args
	hostname, err := os.Hostname()

	for i, arg := range args {
		println("Argument", i, ":", arg)
	}

	if err != nil {
		println("Error getting hostname:", err.Error())
	} else {
		println("Hostname:", hostname)
	}
}
