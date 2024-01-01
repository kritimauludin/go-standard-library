package main

/**
Package strconv

● Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
● Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string,
atau sebaliknya
● Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
● https://golang.org/pkg/strconv/
*/
import (
	"fmt"
	"strconv"
)

func main() {
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	resultInt, err := strconv.Atoi("191991")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)
}
