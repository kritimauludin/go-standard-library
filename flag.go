package main

/**
Package flag

● Package flag berisikan fungsionalitas untuk memparsing command line argument
● https://golang.org/pkg/flag/
*/

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Your database username")
	password := flag.String("password", "", "Your database password")
	host := flag.String("host", "localhost", "Your database host")
	port := flag.Int("port", 8080, "Your database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	// go run flag.go -username=root -password=pass -host=127.0.0.1 -port=5505
}
