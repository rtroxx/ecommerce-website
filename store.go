package main

import (
	"errors"
	"fmt"
)

type Store struct {
	Users     []*User
	Inventory *Inventory
	Orders    []*Order
}

func (s *Store) PlaceOrder(o *Order) (bool, error) {
	for _, u1 := range s.Users {
		if o.User.Name == u1.Name {
			for _, c1 := range o.orderitems {
				if s.Inventory.Availability[c1.Product.Name] < c1.Quantity {
					return false, errors.New("quantity is not available")
				}
			}
			for _, c1 := range o.orderitems {
				s.Inventory.Availability[c1.Product.Name] -= c1.Quantity
			}
			fmt.Println("Order Successfully Placed for users:%s\n", o.User.Name)
		}
	}
	return true, nil
}
