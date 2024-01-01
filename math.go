package main

import (
	"fmt"
	"math"
)

/**
Package math

● Package math merupakan package yang berisikan constant dan fungsi matematika
● https://golang.org/pkg/math/
*/

func main() {
	fmt.Println(math.Ceil(1.50))
	fmt.Println(math.Floor(1.50))
	fmt.Println(math.Round(1.50))
	fmt.Println(math.Min(11, 50))
	fmt.Println(math.Max(11, 50))
}
