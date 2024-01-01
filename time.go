package main

import (
	"fmt"
	"time"
)

/**
Package time

● Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
● https://golang.org/pkg/time/
*/

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2002, time.June, 5, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "02 January 2006 15:04"

	value := "01 January 2022 15:15"
	// value := "asal"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
