package trademodel

import "fmt"

type Currency struct {
	Coin string
	Base string
}

func (currency Currency) String() string {
	return fmt.Sprintf("%s_%s", currency.Coin, currency.Base)
}

type Balance struct {
	Currency  string
	Available float64
	Frozen    float64
	Balance   float64
}
