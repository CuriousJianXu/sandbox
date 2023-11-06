package fn

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	_ "github.com/lib/pq"
	_ "oproaster.com/sandbox/pkg/zloginit"
	"oproaster.com/sandbox/usecase"
)

var (
	uc *usecase.Usecase
)

func init() { //nolint:gochecknoinits
	functions.HTTP("Entry", Entry)

	log.Info().Msg("Connect to CloudSQL")
	log.Info().Str("POSTGRES_HOST", os.Getenv("POSTGRES_HOST")).Str("POSTGRES_USER", os.Getenv("POSTGRES_USER")).Msg("Connect to CloudSQL: Debugging info")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable pool_max_conns=%s pool_min_conns=%s pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s application_name=%s ",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		"2",
		"1",
		"1h",
		"30m",
		"crawler",
	)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Connect to CloudSQL: Failed")
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal().Err(err).Msg("Connect to CloudSQL: Failed")
	}

	uc = usecase.New(pool)
	log.Info().Msg("Connect to CloudSQL: Done. Start serving requests")

}

// Entry is an HTTP Cloud Function with a request parameter.
func Entry(w http.ResponseWriter, r *http.Request) {
	// var d struct {
	// 		Name string `json:"name"`
	// }
	// if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
	// 		fmt.Fprint(w, "Hello, World!")
	// 		return
	// }
	// if d.Name == "" {
	// 		fmt.Fprint(w, "Hello, World!")
	// 		return
	// }
	ctx := context.Background()
	uc.CrawlAndStoreTransactions(ctx)
}
