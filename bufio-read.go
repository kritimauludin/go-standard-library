package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long of string\nthis is short of string\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
