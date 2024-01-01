package main

/**
Package reflect

● Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode
kita pada saat aplikasi sedang berjalan
● Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
● Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package
reflec ini secara otodidak
● Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah
digunakan
● https://golang.org/pkg/reflect/

untuk membaca semua metadata bisa menggunakan reflection ini
*/

import (
	"fmt"
	"reflect"
	"strconv"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Value type", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(strconv.ParseBool(structField.Tag.Get("required")))
		fmt.Println(strconv.Atoi(structField.Tag.Get("max")))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface() //interface adalah dataa
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{"kriti"})
	readField(Person{"mauludin", "bogor", "kriti@gmail.com"})

	person := Person{
		Name:    "kriti",
		Address: "bogor",
		Email:   "ada",
	}

	fmt.Println(isValid(person))
}
