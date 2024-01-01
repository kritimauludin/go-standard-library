package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Uji coba\n")
	_, _ = writer.WriteString("Test\n")

	writer.Flush()
}
