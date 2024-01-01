package main

import (
	"fmt"
	"regexp"
)

/**
Package regexp

● Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
● Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
● https://github.com/google/re2/wiki/Syntax
● https://golang.org/pkg/regexp/
*/

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`k([a-z])i`)

	fmt.Println(regex.MatchString("kri"))
	fmt.Println(regex.MatchString("kri"))
	fmt.Println(regex.MatchString("kRi"))
	fmt.Println(regex.MatchString("kir"))

	fmt.Println(regex.FindAllString("kri kRi kmi krit mam real kti", 10))
}
