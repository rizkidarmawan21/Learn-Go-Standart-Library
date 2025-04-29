package main

import (
	"encoding/csv"
	"os"
)

func main(){
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan"})
	_ = writer.Write([]string{"rizki", "darmawan"})

	// fungsi untuk menyimpan data ke file
	writer.Flush()
}