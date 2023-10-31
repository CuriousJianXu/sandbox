// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package genrepo

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.insertOrdersStmt, err = db.PrepareContext(ctx, InsertOrders); err != nil {
		return nil, fmt.Errorf("error preparing query InsertOrders: %w", err)
	}
	if q.selectItemsStmt, err = db.PrepareContext(ctx, SelectItems); err != nil {
		return nil, fmt.Errorf("error preparing query SelectItems: %w", err)
	}
	if q.selectOrdersByItemIDAndDateStmt, err = db.PrepareContext(ctx, SelectOrdersByItemIDAndDate); err != nil {
		return nil, fmt.Errorf("error preparing query SelectOrdersByItemIDAndDate: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.insertOrdersStmt != nil {
		if cerr := q.insertOrdersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertOrdersStmt: %w", cerr)
		}
	}
	if q.selectItemsStmt != nil {
		if cerr := q.selectItemsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectItemsStmt: %w", cerr)
		}
	}
	if q.selectOrdersByItemIDAndDateStmt != nil {
		if cerr := q.selectOrdersByItemIDAndDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectOrdersByItemIDAndDateStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	insertOrdersStmt                *sql.Stmt
	selectItemsStmt                 *sql.Stmt
	selectOrdersByItemIDAndDateStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		insertOrdersStmt:                q.insertOrdersStmt,
		selectItemsStmt:                 q.selectItemsStmt,
		selectOrdersByItemIDAndDateStmt: q.selectOrdersByItemIDAndDateStmt,
	}
}
