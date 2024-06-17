package server

import "time"

type Item struct {
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	PriceDate time.Time `json:"pricedate"`
}
