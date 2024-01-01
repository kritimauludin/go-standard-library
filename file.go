package main

/**
File Management

● Di package os, terdapat File Management, namun sengaja ditunda pembahasannya, karena kita
harus tahu dulu tentang IO
● Saat kita membuat atau membaca file menggunakan Package os, struct File merupakan
implementasi dari io.Reader dan io.Writer
● Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package
io / bufio

Open File

● Untuk membuat / membaca File, kita bisa menggunakan os.OpenFile(name, flag, permission)
● name berisikan nama file, bisa absolute atau relative / local
● flag merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain
● permission, merupakan permission yang diperlukan ketika membuat file, bisa kita simulasikan
disini : https://chmod-calculator.com/
*/
import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil

}
func main() {
	// createNewFile("sample.log", "This is simple log")

	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	addToFile("sample.log", "\nThis add new log")
}
