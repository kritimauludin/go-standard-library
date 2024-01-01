package main

/**
Package strings

● Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data
String
● Ada banyak sekali function yang bisa kita gunakan
● https://golang.org/pkg/strings/
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Kriti Mauludin", "kriti"))
	fmt.Println(strings.Split("Kriti Mauludin", " "))
	fmt.Println(strings.ToLower("Kriti Mauludin"))
	fmt.Println(strings.ToUpper("Kriti Mauludin"))
	fmt.Println(strings.Trim("        Kriti Mauludin       ", " "))
	fmt.Println(strings.ReplaceAll("Kriti Mauludin Kriti Maul", "Kriti", "Aul"))
}
