package repo

import (
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
