package main

import (
	"fmt"
	"time"
)

/**
Duration

● Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
● Package tipe memiliki type Duration, yang sebenarnya adalah alias untuk int64
● Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration
*/

func main() {
	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration : %d\n", duration3)
}
