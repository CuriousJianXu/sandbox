// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package genrepo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Item struct {
	ID   int32       `db:"id"`
	Name pgtype.Text `db:"name"`
}

type Order struct {
	ID     int32  `db:"id"`
	Date   string `db:"date"`
	ItemID int32  `db:"item_id"`
	Count  int32  `db:"count"`
	Price  int32  `db:"price"`
}
