package main

/**
Package errors

● Sebelumnya kita sudah membahas tentang interface error yang merupakan representasi dari error
di Go-Lang, dan membuat error menggunakan function errors.New()
● Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya ketika
kita ingin membuat beberapa value error yang berbeda
● https://pkg.go.dev/errors

Mengecek Jenis Error

● Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
● Kita bisa menggunakan errors.Is() untuk mengecek jenis type error nya
*/

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "kriti" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("kriti")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error!")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error!")
		} else {
			fmt.Println("Unknown error!")
		}
	}
}
