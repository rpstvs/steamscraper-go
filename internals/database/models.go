// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"
)

type Item struct {
	ID       int32
	Itemname string
}

type Price struct {
	Pricedate time.Time
	ItemID    int32
	Price     string
}
