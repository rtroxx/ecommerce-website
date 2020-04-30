package main

//package main
type Order struct {
	User            User
	orderitems      []*OrderItem
	ShippingAddress string
}

func (o *Order) TotalWithoutTax() int {
	cost := 0
	for _, c1 := range o.orderitems {
		cost += c1.Product.Price * c1.Quantity
	}
	return cost
}
func (o *Order) TotalWithTax() int {
	cost := 0
	for _, c1 := range o.orderitems {
		cost += c1.Total
	}
	return cost
}
