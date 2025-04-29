package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main(){
	csvString := "eko,kurniawan\n" + 
		"rizki,darmawan\n" +
		"john,doe\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}