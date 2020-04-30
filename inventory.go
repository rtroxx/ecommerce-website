package main

type Inventory struct {
	Products     []*Product
	Availability map[string]int
}

func NewInvetory() *Inventory {
	return &Inventory{
		Availability: make(map[string]int),
	}
}
func (s *Inventory) AddProduct(p *Product) {
	if s.Availability == nil {
		s.Availability = make(map[string]int)
	}
	if s.Availability[p.Name] == 0 {
		s.Products = append(s.Products, p)
		s.Availability[p.Name]++
	} else {
		s.Availability[p.Name]++
	}

}
func (s *Inventory) RemoveProduct(p *Product) {
	var ret []*Product
	for _, curr := range s.Products {
		if curr != p {
			ret = append(ret, curr)
		}
	}
	if s.Availability[p.Name] != 0 {
		s.Availability[p.Name]--
	}
}
func (s *Inventory) Quantity(p *Product) int {
	return s.Availability[p.Name]
}
