package main

import "fmt"

//Product => Name, Categories (multiple), Price, Description
type Product struct {
	Name        string
	Categories  []string
	Price       int
	Description string
}

func NewProduct(name string, categories []string, p int, des string) *Product {
	return &Product{
		Name:        name,
		Categories:  append([]string{}, categories...),
		Price:       p,
		Description: des,
	}
}
func (u *Product) show() {
	fmt.Println(u.Name)
	fmt.Println(u.Categories)
	fmt.Println(u.Price)
	fmt.Printf(u.Description)
}
