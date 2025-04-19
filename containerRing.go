// Container Ring adalah package yang berisi fungsi-fungsi untuk mengelola ring
// Implementasi struktur data circular list
// Circular list adalah list yang memiliki pointer ke elemen pertama dan terakhir
// Dimana akhir element akan kembali ke awal element (HEAD)

package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main(){
	data := ring.New(5) // Membuat ring baru dengan kapasitas 5
	data2 := ring.New(5) // Membuat ring baru dengan kapasitas 5

	data.Value = 1 // Menambahkan nilai 1 ke ring

	data = data.Next() // Pindah ke elemen selanjutnya
	data.Value = 2 // Menambahkan nilai 2 ke ring

	data = data.Next() // Pindah ke elemen selanjutnya
	data.Value = 3 // Menambahkan nilai 3 ke ring

	data = data.Next() // Pindah ke elemen selanjutnya
	data.Value = 4 // Menambahkan nilai 4 ke ring

	data = data.Next() // Pindah ke elemen selanjutnya
	data.Value = 5 // Menambahkan nilai 5 ke ring


	// sekarang kalau mau otomatisasi pakai for loop

	for i := 0; i < data2.Len(); i++ {
		data2.Value = "Value " + strconv.Itoa(i)
		data2 = data2.Next() // Pindah ke elemen selanjutnya / next
	}

	// jika ingin melakukan iterasi dari tiap datanya
	data.Do(func(value interface{}) {
		println(value.(int))
	})

	data2.Do(func(value interface{}) {
		fmt.Println(value)
	})

}