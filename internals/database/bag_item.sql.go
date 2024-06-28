// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: bag_item.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const addItem = `-- name: AddItem :one
INSERT INTO Bag_item (Bag_id, Item_id, Amount)
VALUES ($1, $2, $3)
RETURNING bag_id, item_id, amount
`

type AddItemParams struct {
	BagID  uuid.UUID
	ItemID uuid.UUID
	Amount int32
}

func (q *Queries) AddItem(ctx context.Context, arg AddItemParams) (BagItem, error) {
	row := q.db.QueryRowContext(ctx, addItem, arg.BagID, arg.ItemID, arg.Amount)
	var i BagItem
	err := row.Scan(&i.BagID, &i.ItemID, &i.Amount)
	return i, err
}

const deleteItem = `-- name: DeleteItem :one
DELETE FROM Bag_item
WHERE Bag_id = $1
    AND Item_id = $2
RETURNING bag_id, item_id, amount
`

type DeleteItemParams struct {
	BagID  uuid.UUID
	ItemID uuid.UUID
}

func (q *Queries) DeleteItem(ctx context.Context, arg DeleteItemParams) (BagItem, error) {
	row := q.db.QueryRowContext(ctx, deleteItem, arg.BagID, arg.ItemID)
	var i BagItem
	err := row.Scan(&i.BagID, &i.ItemID, &i.Amount)
	return i, err
}

const getBagItem = `-- name: GetBagItem :one
SELECT Amount
From Bag_item
WHERE BAG_id = $1
    AND Item_id = $2
`

type GetBagItemParams struct {
	BagID  uuid.UUID
	ItemID uuid.UUID
}

func (q *Queries) GetBagItem(ctx context.Context, arg GetBagItemParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getBagItem, arg.BagID, arg.ItemID)
	var amount int32
	err := row.Scan(&amount)
	return amount, err
}

const updateBagItem = `-- name: UpdateBagItem :one
UPDATE Bag_item
SET Amount = $2
WHERE Bag_id = $1
    AND Item_id = $3
RETURNING bag_id, item_id, amount
`

type UpdateBagItemParams struct {
	BagID  uuid.UUID
	Amount int32
	ItemID uuid.UUID
}

func (q *Queries) UpdateBagItem(ctx context.Context, arg UpdateBagItemParams) (BagItem, error) {
	row := q.db.QueryRowContext(ctx, updateBagItem, arg.BagID, arg.Amount, arg.ItemID)
	var i BagItem
	err := row.Scan(&i.BagID, &i.ItemID, &i.Amount)
	return i, err
}
