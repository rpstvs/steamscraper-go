// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: item.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createItem = `-- name: CreateItem :one
INSERT INTO Items (id, ItemName, ImageUrl, DayChange, WeekChange)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, itemname, classid, daychange, weekchange, imageurl
`

type CreateItemParams struct {
	ID         uuid.UUID
	Itemname   string
	Imageurl   string
	Daychange  float64
	Weekchange float64
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.ID,
		arg.Itemname,
		arg.Imageurl,
		arg.Daychange,
		arg.Weekchange,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Itemname,
		&i.Classid,
		&i.Daychange,
		&i.Weekchange,
		&i.Imageurl,
	)
	return i, err
}

const getItemByName = `-- name: GetItemByName :one
SELECT Id
FROM Items
WHERE itemname = $1
`

func (q *Queries) GetItemByName(ctx context.Context, itemname string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getItemByName, itemname)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getItemsIds = `-- name: GetItemsIds :many
SELECT Id
FROM Items
`

func (q *Queries) GetItemsIds(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getItemsIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDailyChange = `-- name: UpdateDailyChange :exec
UPDATE Items
SET DayChange = $1
WHERE Id = $2
`

type UpdateDailyChangeParams struct {
	Daychange float64
	ID        uuid.UUID
}

func (q *Queries) UpdateDailyChange(ctx context.Context, arg UpdateDailyChangeParams) error {
	_, err := q.db.ExecContext(ctx, updateDailyChange, arg.Daychange, arg.ID)
	return err
}

const updateWeeklyChange = `-- name: UpdateWeeklyChange :exec
UPDATE Items
SET WeekChange = $1
WHERE Id = $2
`

type UpdateWeeklyChangeParams struct {
	Weekchange float64
	ID         uuid.UUID
}

func (q *Queries) UpdateWeeklyChange(ctx context.Context, arg UpdateWeeklyChangeParams) error {
	_, err := q.db.ExecContext(ctx, updateWeeklyChange, arg.Weekchange, arg.ID)
	return err
}
