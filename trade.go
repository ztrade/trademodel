package trademodel

import "time"

const (
	ActionBuy = iota
	ActionSell
	ActionOpenLong
	ActionCloseLong
	ActionOpenShort
	ActionCloseShort
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

func (ta *TradeAction) IsBuy() bool {
	switch ta.Action {
	case ActionBuy, ActionOpenLong, ActionCloseShort:
		return true
	case ActionSell, ActionCloseLong, ActionOpenShort:
		return false
	default:
	}
	return false
}
