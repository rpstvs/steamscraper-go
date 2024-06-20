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

const getLatestPrice = `-- name: GetLatestPrice :one
SELECT Price,
    PriceDate
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC
`

type GetLatestPriceRow struct {
	Price     float64
	Pricedate time.Time
}

func (q *Queries) GetLatestPrice(ctx context.Context, itemID uuid.UUID) (GetLatestPriceRow, error) {
	row := q.db.QueryRowContext(ctx, getLatestPrice, itemID)
	var i GetLatestPriceRow
	err := row.Scan(&i.Price, &i.Pricedate)
	return i, err
}

const getPricebyId = `-- name: GetPricebyId :many
SELECT Price,
    PriceDate
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC
`

type GetPricebyIdRow struct {
	Price     float64
	Pricedate time.Time
}

func (q *Queries) GetPricebyId(ctx context.Context, itemID uuid.UUID) ([]GetPricebyIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getPricebyId, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPricebyIdRow
	for rows.Next() {
		var i GetPricebyIdRow
		if err := rows.Scan(&i.Price, &i.Pricedate); err != nil {
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
