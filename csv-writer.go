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
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"kriti", "21", "Bogor"})
	_ = writer.Write([]string{"indah", "21", "Bogor"})
	_ = writer.Write([]string{"mauludin", "21", "Bogor"})

	writer.Flush()
}
