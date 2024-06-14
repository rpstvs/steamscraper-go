// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: price.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addPrice = `-- name: AddPrice :many
INSERT INTO Prices (PriceDate, Item_id, Price)
VALUES ($1, $2, $3)
RETURNING pricedate, item_id, price
`

type AddPriceParams struct {
	Pricedate time.Time
	ItemID    uuid.UUID
	Price     float64
}

func (q *Queries) AddPrice(ctx context.Context, arg AddPriceParams) ([]Price, error) {
	rows, err := q.db.QueryContext(ctx, addPrice, arg.Pricedate, arg.ItemID, arg.Price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Price
	for rows.Next() {
		var i Price
		if err := rows.Scan(&i.Pricedate, &i.ItemID, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPricebyId = `-- name: GetPricebyId :one
SELECT Price
FROM Prices
WHERE Item_id = $1
`

func (q *Queries) GetPricebyId(ctx context.Context, itemID uuid.UUID) (float64, error) {
	row := q.db.QueryRowContext(ctx, getPricebyId, itemID)
	var price float64
	err := row.Scan(&price)
	return price, err
}
