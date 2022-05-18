package trademodel

import (
	"fmt"
	"time"
)

type TradeType int

const (
	CancelOne   TradeType = -2
	CancelAll   TradeType = -1
	DirectLong  TradeType = 1
	DirectShort TradeType = 1 << 1

	Limit  TradeType = 1 << 3
	Market TradeType = 1 << 4
	Stop   TradeType = 1 << 5

	Open  TradeType = 1 << 6
	Close TradeType = 1 << 7

	OpenLong   = Open | DirectLong
	OpenShort  = Open | DirectShort
	CloseLong  = Close | DirectLong
	CloseShort = Close | DirectShort
	StopLong   = Stop | DirectLong
	StopShort  = Stop | DirectShort
)

func (t TradeType) String() (ret string) {
	if t&Limit == Limit {
		ret += "Limit"
	} else if t&Market == Market {
		ret += "Market"
	} else if t&Stop == Stop {
		ret += "Stop"
	}
	if t&Open == Open {
		ret += "Open"
	} else if t&Close == Close {
		ret += "Close"
	}

	if t&DirectLong == DirectLong {
		ret += "Long"
	} else if t&DirectShort == DirectShort {
		ret += "Short"
	}
	return
}

func NewTradeType(typ string) (t TradeType, err error) {
	switch typ {
	case "OpenLong":
		t = OpenLong
	case "OpenShort":
		t = OpenShort
	case "CloseLong":
		t = CloseLong
	case "CloseShort":
		t = CloseShort
	case "StopLong":
		t = StopLong
	case "StopShort":
		t = StopShort
	default:
		err = fmt.Errorf("unsupport trade type: %d", t)
	}
	return
}

type Trade struct {
	ID     string
	Action TradeType
	Time   time.Time
	Price  float64
	Amount float64
	Side   string
	Remark string
}

// TradeAction trade action
type TradeAction struct {
	ID     string
	Action TradeType
	Amount float64
	Price  float64
	Time   time.Time
	Symbol string
}

func (a TradeType) IsLong() bool {
	if a&OpenLong == OpenLong || a&CloseShort == CloseShort || a&StopShort == StopShort {
		return true
	}
	return false
}

func (a TradeType) IsOpen() bool {
	if a&Open == Open {
		return true
	}
	return false
}

func (a TradeType) IsStop() bool {
	if a&Stop == Stop {
		return true
	}
	return false
}
