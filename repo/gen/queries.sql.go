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
	_, err := q.exec(ctx, q.insertOrdersStmt, InsertOrders,
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

func (q *Queries) SelectItems(ctx context.Context) ([]Item, error) {
	rows, err := q.query(ctx, q.selectItemsStmt, SelectItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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
