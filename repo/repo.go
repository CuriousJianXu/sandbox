package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	genrepo "oproaster.com/sandbox/repo/gen"
)

type Repo struct {
	q *genrepo.Queries
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		q: genrepo.New(db),
	}
}
