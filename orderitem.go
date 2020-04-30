package main

type OrderItem struct {
	Product  *Product
	Quantity int
	Tax      int
	Total    int
}

func NewOrderItem(p *Product, q int) *OrderItem {
	return &OrderItem{
		Product:  p,
		Quantity: q,
		Tax:      10,
	}

}
