// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID
	Itemname string
}

type Price struct {
	Pricedate time.Time
	ItemID    uuid.UUID
	Price     float64
}

type User struct {
	ID  uuid.UUID
	Bag uuid.NullUUID
}
