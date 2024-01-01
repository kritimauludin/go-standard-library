package main

/**
Package encoding

● Golang menyediakan package encoding untuk melakukan encode dan decode
● https://pkg.go.dev/encoding
● Golang menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah
base64, csv dan json
*/

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Kriti, 21, Bogor\n" +
		"indah, 21, Bogor\n" +
		"mauludin,  21, Bogor"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
