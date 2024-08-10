// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Bag struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Totalvalue float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
}

type BagItem struct {
	BagID  uuid.UUID
	ItemID uuid.UUID
	Amount int32
}

type Item struct {
	ID         uuid.UUID
	Itemname   string
	Daychange  float64
	Weekchange float64
	Imageurl   string
}

type Price struct {
	Pricedate time.Time
	ItemID    uuid.UUID
	Price     float64
}

type User struct {
	ID        uuid.UUID
	Name      string
	Steamid   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
