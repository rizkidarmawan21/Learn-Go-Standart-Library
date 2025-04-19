// Package Math adalah package yang berisi fungsi-fungsi untuk melakukan operasi matematika

package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(math.Ceil(1.5)) // 2: Mengembalikan nilai bulat terkecil yang lebih besar atau sama dengan 1.5
	fmt.Println(math.Floor(1.5)) // 1: Mengembalikan nilai bulat terbesar yang lebih kecil atau sama dengan 1.5
	fmt.Println(math.Abs(-1.5)) // 1.5: Mengembalikan nilai absolut dari -1.5
	fmt.Println(math.Pow(2, 3)) // 8: Mengembalikan 2 pangkat 3 (2^3)
	fmt.Println(math.Sqrt(16)) // 4: Mengembalikan akar kuadrat dari 16
	fmt.Println(math.Max(1, 2)) // 2: Mengembalikan nilai maksimum antara 1 dan 2
	fmt.Println(math.Min(1, 2)) // 1: Mengembalikan nilai minimum antara 1 dan 2
	fmt.Println(math.Round(1.5)) // 2: Mengembalikan nilai 1.5 dibulatkan ke nilai bulat terdekat	
}