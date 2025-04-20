// Package time adalah package yang berisi fungsi-fungsi untuk mengelola waktu
// Now untuk mendapatkan waktu saat ini
// Parse untuk mengubah string menjadi waktu
// Date untuk membuat waktu dari tahun, bulan, hari, jam, menit, detik

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current Time:", now)

	utc := time.Date(2009, time.November ,10, 23, 0, 0, 0, time.UTC)
	fmt.Println("UTC Time:", utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println("Parsed Time:", parse)
}