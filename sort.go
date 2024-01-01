package main

/**
Package sort

● Package sort adalah package yang berisikan utilitas untuk proses pengurutan
● Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
● https://golang.org/pkg/sort/
*/
import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// slice defaultnya pointer
type UserSlice []User

func (a UserSlice) Len() int           { return len(a) }
func (a UserSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserSlice) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	users := []User{
		{"Kriti", 15},
		{"Mauludin", 17},
		{"Aul", 20},
		{"Maul", 18},
		{"Krit", 22},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
