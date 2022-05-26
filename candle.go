package trademodel

import (
	"fmt"
	"time"
)

type Candle struct {
	ID       int64   `xorm:"pk autoincr null 'id'"`
	Start    int64   `xorm:"unique index 'start'"`
	Open     float64 `xorm:"notnull 'open'"`
	High     float64 `xorm:"notnull 'high'"`
	Low      float64 `xorm:"notnull 'low'"`
	Close    float64 `xorm:"notnull 'close'"`
	Volume   float64 `xorm:"notnull 'volume'"`
	Turnover float64 `xorm:"turnover 'turnover'"`
	Trades   int64   `xorm:"notnull 'trades'"`
	Table    string  `xorm:"-"`
}

func (c Candle) TableName() string {
	return c.Table
}

func (c Candle) GetTable() string {
	return c.Table
}

func (c Candle) GetStart() int64 {
	return c.Start
}

func (c *Candle) SetTable(tbl string) {
	c.Table = tbl
}

func (c Candle) Time() time.Time {
	return time.Unix(c.Start, 0)
}

func (c Candle) String() string {
	return fmt.Sprintf("%s open:%f close:%f low:%f high:%f volume:%f trades:%d turnover: %f", c.Time().String(), c.Open, c.Close, c.Low, c.High, c.Volume, c.Trades, c.Turnover)
}

// CandleList candle list
type CandleList []*Candle

// Merge merge multi candle to one
func (l CandleList) Merge() (ret *Candle) {
	if len(l) == 0 {
		return
	}
	ret = new(Candle)
	ret.Start = l[0].Start
	ret.Open = l[0].Open
	ret.High = l.High()
	ret.Low = l.Low()
	ret.Close = l[len(l)-1].Close
	for _, v := range l {
		ret.Turnover = v.Turnover
		ret.Volume += v.Volume
		ret.Trades += v.Trades
	}
	return
}

func (l CandleList) High() (ret float64) {
	for _, v := range l {
		if ret < v.High {
			ret = v.High
		}
	}
	return
}

func (l CandleList) Low() (ret float64) {
	for _, v := range l {
		if ret == 0 {
			ret = v.Low
			continue
		}
		if ret > v.Low {
			ret = v.Low
		}
	}
	return
}
