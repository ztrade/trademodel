package trademodel

import (
	"strings"
	"time"
)

const (
	SymbolTypeSpot    = "spot"
	SymbolTypeIndex   = "index"
	SymbolTypeFutures = "futures"
)

type Symbol struct {
	Name        string `json:"name"`
	Symbol      string `json:"code"`
	Exchange    string `json:"exchange"`
	Description string `json:"description"`
	Type        string `json:"type"`
	// 交易时间
	Session string `json:"session"`

	// price decimal places
	Precision int `json:"precision"`
	// amount decimal places
	AmountPrecision int `json:"volume_precision"`

	Expired bool `json:"expired"`
	// only use if expired = true
	ExpirationDate time.Time `json:"expiration_date"`
	Resolutions    string    `json:"resolutions"`

	// mini size between two price
	PriceStep float64 `json:"price_step"`
	// mini size between two amount
	AmountStep float64 `json:"amount"`
}

func (s *Symbol) GetResolutions() []string {
	return strings.Split(s.Resolutions, ",")
}

func (s *Symbol) FixPrice(price float64) float64 {
	return float64(int(price/s.PriceStep)) * s.PriceStep
}
