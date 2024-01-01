package main

import (
	"container/list"
	"fmt"
)

/**
Package container/list

● Package container/list adalah implementasi struktur data double linked list di Go-Lang
● https://golang.org/pkg/container/list/
*/

func main() {
	var data *list.List = list.New()

	data.PushBack("Kriti")
	data.PushBack("Mauludin")
	data.PushBack("Bogor")
	data.PushBack("Barat")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // kriti

	next := head.Next()
	fmt.Println(next.Value) //Mauludin

	next = next.Next()
	fmt.Println(next.Value) //bogor

	//menggunakan for loop
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
