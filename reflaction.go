// reflact adalah package yang berisi fungsi-fungsi untuk merefleksi data
// dimana kita dapat melihat struktur kode kita pada aplikasi yang sedang berjalan
// ini cocok untuk kita ketika jika ingin membuat aplikasi yang dinamis/general

package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

// StructTag ini cocok untuk validasi
type Person struct {
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func readField(value any){
	valueType := reflect.TypeOf(value)

	fmt.Println("Type:", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Println("Field Name:", field.Name)
		fmt.Println("Field Type:", field.Type)

		fieldTag := field.Tag
		fmt.Println("Field Tag Required:", fieldTag.Get("required"))
		fmt.Println("Field Tag Max:", fieldTag.Get("max"))
	}
}

func IsValid(data any) (result bool){
	tag := reflect.TypeOf(data)
	for i := 0; i < tag.NumField(); i++ {
		field := tag.Field(i)
		if field.Tag.Get("required") == "true" {
			dataValue := reflect.ValueOf(data).Field(i).Interface()
			result = dataValue != ""
			if !result {
				fmt.Println("Field", field.Name, "is required!")
				return false
			}
		}
	}

	return true
}

func main(){
	readField(Sample{"Eko"})
	readField(Person{"Eko2","",""})

	// validasi
	person := Person{
		Name:    "Eko",
		Address: "Jakarta",
		Email:   "eko@gmail.com",
	}

	fmt.Println(IsValid(person))
}