package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	genrepo "oproaster.com/sandbox/repo/gen"
)

type Repo struct {
	q *genrepo.Queries
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		q: genrepo.New(db),
	}
}

func (r *Repo) InsertOrders(ctx context.Context, arg genrepo.InsertOrdersParams) error {
	return r.q.InsertOrders(ctx, arg)
}
