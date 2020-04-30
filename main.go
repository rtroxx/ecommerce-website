package main

import "fmt"

func main() {
	//u1 := NewUser("Ratnesh", 18, "Rajiv Nagar")
	//c := Inventory{}
	x := NewInvetory()
	p1 := NewProduct("cow", []string{"r1", "r2"}, 1500, "very good")
	p2 := NewProduct("dog", []string{"s1", "s2"}, 2000, "good")
	x.AddProduct(p1)
	x.AddProduct(p2)
	fmt.Println(x.Quantity(p1)) // ==> Should print 2.
	x.RemoveProduct(p1)
	fmt.Println(x.Quantity(p2)) //=> Should print 1.

}
