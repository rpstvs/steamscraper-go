// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: bag.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBag = `-- name: CreateBag :one
INSERT INTO Bag (Id, TotalValue, User_id, Created_at, Updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, totalvalue, created_at, updated_at
`

type CreateBagParams struct {
	ID         uuid.UUID
	Totalvalue float64
	UserID     uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (q *Queries) CreateBag(ctx context.Context, arg CreateBagParams) (Bag, error) {
	row := q.db.QueryRowContext(ctx, createBag,
		arg.ID,
		arg.Totalvalue,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Bag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Totalvalue,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBagbyID = `-- name: GetBagbyID :one
SELECT id, user_id, totalvalue, created_at, updated_at
FROM Bag
WHERE Id = $1
`

func (q *Queries) GetBagbyID(ctx context.Context, id uuid.UUID) (Bag, error) {
	row := q.db.QueryRowContext(ctx, getBagbyID, id)
	var i Bag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Totalvalue,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBagsByUser = `-- name: GetBagsByUser :many
SELECT id, user_id, totalvalue, created_at, updated_at
FROM Bag
Where User_id = $1
`

func (q *Queries) GetBagsByUser(ctx context.Context, userID uuid.UUID) ([]Bag, error) {
	rows, err := q.db.QueryContext(ctx, getBagsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bag
	for rows.Next() {
		var i Bag
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Totalvalue,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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

const updateBag = `-- name: UpdateBag :one
UPDATE Bag
SET TotalValue = $2
    AND Updated_at = $3
WHERE Id = $1
RETURNING id, user_id, totalvalue, created_at, updated_at
`

type UpdateBagParams struct {
	ID         uuid.UUID
	Totalvalue float64
	UpdatedAt  time.Time
}

func (q *Queries) UpdateBag(ctx context.Context, arg UpdateBagParams) (Bag, error) {
	row := q.db.QueryRowContext(ctx, updateBag, arg.ID, arg.Totalvalue, arg.UpdatedAt)
	var i Bag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Totalvalue,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
