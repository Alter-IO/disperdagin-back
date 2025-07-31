package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type CommodityItem struct {
	ID    string          `json:"id"`
	Price decimal.Decimal `json:"price"`
}

type CommodityDaily struct {
	Commodities []CommodityItem `json:"commodities"`
	PublishDate time.Time       `json:"publish_date"`
}
