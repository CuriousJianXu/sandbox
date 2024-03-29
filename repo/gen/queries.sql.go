// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: queries.sql

package genrepo

import (
	"context"
)

const InsertOrders = `-- name: InsertOrders :exec
INSERT INTO orders (date, item_id, count, price) 
VALUES (
    $1,
    $2,
    $3,
    $4
)
`

type InsertOrdersParams struct {
	Date   string `db:"date"`
	ItemID int32  `db:"item_id"`
	Count  int32  `db:"count"`
	Price  int32  `db:"price"`
}

func (q *Queries) InsertOrders(ctx context.Context, arg InsertOrdersParams) error {
	_, err := q.db.Exec(ctx, InsertOrders,
		arg.Date,
		arg.ItemID,
		arg.Count,
		arg.Price,
	)
	return err
}

const SelectItems = `-- name: SelectItems :many
SELECT id, name FROM items
ORDER BY id
`

func (q *Queries) SelectItems(ctx context.Context) ([]*Item, error) {
	rows, err := q.db.Query(ctx, SelectItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const SelectOrdersByItemIDAndDate = `-- name: SelectOrdersByItemIDAndDate :many
SELECT id, date, item_id, count, price FROM orders
WHERE item_id=$1
  AND date=$2
`

type SelectOrdersByItemIDAndDateParams struct {
	ItemID int32  `db:"item_id"`
	Date   string `db:"date"`
}

func (q *Queries) SelectOrdersByItemIDAndDate(ctx context.Context, arg SelectOrdersByItemIDAndDateParams) ([]*Order, error) {
	rows, err := q.db.Query(ctx, SelectOrdersByItemIDAndDate, arg.ItemID, arg.Date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.ItemID,
			&i.Count,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
