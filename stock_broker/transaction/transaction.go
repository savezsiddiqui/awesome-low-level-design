package transaction

import (
	"time"

	"example.com/stock_broker/stock"
)

type OrderType string

const (
	Buy  OrderType = "buy"
	Sell OrderType = "sell"
)

type Transaction struct {
	Timestamp                     time.Time
	Stock                         stock.Stock
	StockPriceatTimeofTransaction float64
	Qty                           int
	OrderType                     OrderType
}
