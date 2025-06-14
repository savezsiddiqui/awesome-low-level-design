package stock

type Stock struct {
	name  string
	price float64
}

func NewStock(name string, price float64) *Stock {
	return &Stock{
		name:  name,
		price: price,
	}
}

func (s *Stock) UpdatePrice(newPrice float64) {
	s.price = newPrice
}

func (s *Stock) GetPrice() float64 {
	return s.price
}

func (s *Stock) GetName() string {
	return s.name
}
