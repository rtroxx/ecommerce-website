package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address string
}

func NewUser(name string, age int, address string) *User {
	return &User{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

//Name, Age, Address.
func (u *User) show() {
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	fmt.Println(u.Address)
}
