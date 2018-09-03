package trademodel

import "time"

const (
	ActionBuy  = 1
	ActionSell = 2
)

type Trade struct {
	ID     string
	Time   time.Time
	Price  float64
	Amount float64
	Side   string
	Remark string
}

type TradeAction struct {
	Action int
	Amount float64
	Price  float64
}
