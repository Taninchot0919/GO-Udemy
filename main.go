package main

import (
	"fmt"
	"time"
)

type User struct{
	Firstname string
	Lastname string
	PhoneNumber string
	Age int
	BirthDate time.Time
}
func main() {
	user := User{
		Firstname: "Taninchot",
		Lastname: "Phuwaloertthiwat",
		PhoneNumber: "0950385146",
		Age: 20,
	}
	fmt.Println(user.Firstname,user.Lastname)
}