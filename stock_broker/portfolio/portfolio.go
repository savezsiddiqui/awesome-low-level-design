package portfolio

import (
	custom_error "example.com/stock_broker/error"
	"example.com/stock_broker/stock"
)

type Portfolio struct {
	holdings map[*stock.Stock]int
}

func NewPortfolio() Portfolio {
	return Portfolio{
		holdings: make(map[*stock.Stock]int),
	}
}

func (p *Portfolio) AddStock(stock *stock.Stock, qty int) {
	_, ok := p.holdings[stock]
	if ok {
		p.holdings[stock] += qty
		return
	}

	p.holdings[stock] = qty
}

func (p *Portfolio) RemoveStock(stock *stock.Stock, qty int) error {
	currentQty, ok := p.holdings[stock]
	if !ok || currentQty < qty {
		return custom_error.NewInsufficientStockError()
	}

	p.holdings[stock] -= qty
	if p.holdings[stock] == 0 {
		delete(p.holdings, stock)
	}
	return nil
}
