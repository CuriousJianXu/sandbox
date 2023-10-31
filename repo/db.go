package repo

import (
	"context"

	genrepo "oproaster.com/sandbox/repo/gen"
)

func (r *Repo) InsertOrders(ctx context.Context, arg genrepo.InsertOrdersParams) error {
	return r.q.InsertOrders(ctx, arg)
}

func (r *Repo) SelectItems(ctx context.Context) ([]genrepo.Item, error) {
	return r.q.SelectItems(ctx)
}

func (r *Repo) SelectOrdersByItemIDAndDate(ctx context.Context, arg genrepo.SelectOrdersByItemIDAndDateParams) ([]genrepo.Order, error) {
	return r.q.SelectOrdersByItemIDAndDate(ctx, arg)
}
