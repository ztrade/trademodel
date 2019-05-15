package trademodel

import "time"

type Order struct {
	OrderID  string
	Symbol   string
	Currency string
	Amount   float64
	Price    float64
	Status   string
	Side     string
	Time     time.Time
}
