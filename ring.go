package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/**
Package container/ring

● Package container/ring adalah implementasi struktur data circular list
● Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal
(HEAD)
● https://golang.org/pkg/container/ring/
*/

func main() {
	var data *ring.Ring = ring.New(5)

	// data.Value = 1

	// data = data.Next()
	// data.Value = 2

	// data = data.Next()
	// data.Value = "3"

	// data = data.Next()
	// data.Value = 4

	// data = data.Next()
	// data.Value = 5

	for i := 1; i <= 5; i++ {
		data.Value = "Value" + strconv.Itoa(i)

		data = data.Next()
	}
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
