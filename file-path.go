package main

/**
Package path

● Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File
System
● Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok
untuk data URL
● https://pkg.go.dev/path
● Namun jika ingin menggunakan untuk memanipulasi path di File System, karena Windows
menggunakan backslash, maka khusus untuk File System, perlu menggunakan pacakge
path/filepath
● https://pkg.go.dev/path/filepath
*/

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go")) //dimulai dari depan mical c:/
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
