package binance

import (
	"time"
)

type Transaction struct {
	Date     time.Time
	Pair     string
	Side     string
	Price    float64
	Executed string
	Amount   string
	Fee      string
}
