package main

import (
	"context"

	"github.com/rs/zerolog/log"
	_ "oproaster.com/sandbox/pkg/zloginit"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"oproaster.com/sandbox/usecase"
)

func main() {
	connectionStr := "" // hidden for security reasons
	db, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	uc := usecase.New(db)

	ctx := context.Background()
	uc.CrawlAndStoreTransactions(ctx)
}
