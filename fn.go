package fn

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "oproaster.com/sandbox/pkg/zloginit"
	"oproaster.com/sandbox/usecase"
)

var (
	uc *usecase.Usecase
)

func init() { //nolint:gochecknoinits
	functions.HTTP("Entry", Entry)

	connectionStr := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"))

	db, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	uc = usecase.New(db)
	log.Info().Msg("Server ready")

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
