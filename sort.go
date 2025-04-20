// Sort adalah package yang berisi fungsi-fungsi untuk mengurutkan data

package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	// temp := u[i]
	// u[i] = u[j]
	// u[j] = temp

	u[i], u[j] = u[j], u[i]
}

func main(){
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}