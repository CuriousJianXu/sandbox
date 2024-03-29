package usecase

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"oproaster.com/sandbox/dto"
	"oproaster.com/sandbox/repo"
	genrepo "oproaster.com/sandbox/repo/gen"
	"oproaster.com/sandbox/utils"
)

type Usecase struct {
	Repo *repo.Repo
}

func New(db *pgxpool.Pool) *Usecase {
	return &Usecase{
		Repo: repo.New(db),
	}
}
func (uc *Usecase) CrawlAndStoreTransactions(ctx context.Context) {
	items, err := uc.Repo.SelectItems(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting items")
		return
	}
	for _, item := range items {
		// if int(item.ID) != 19162 {
		// 	continue
		// }
		// for daysBefore := 1; daysBefore <= 5; daysBefore++ {
		// 	date, err := utils.GetLocalTimeInTaipei(daysBefore)
		// 	if err != nil {
		// 		log.Error().Err(err).Msg("Error getting local date yesterday")
		// 		return
		// 	}
		// 	uc.CrawlAndStoreTransactionsPerItem(ctx, int(item.ID), item.Name.String, date)
		// }
		date, err := utils.GetLocalTimeInTaipei(1)
		if err != nil {
			log.Error().Err(err).Msg("Error getting local date yesterday")
			return
		}
		uc.CrawlAndStoreTransactionsPerItem(ctx, int(item.ID), item.Name.String, date)
	}
}

func (uc *Usecase) CrawlAndStoreTransactionsPerItem(ctx context.Context, itemID int, itemName string, date string) {

	ordersOfYesterday, err := uc.Repo.SelectOrdersByItemIDAndDate(ctx, genrepo.SelectOrdersByItemIDAndDateParams{ItemID: int32(itemID), Date: date})
	if err != nil {
		log.Error().Err(err).Int("item_id", itemID).Str("item_name", itemName).Msg("Error getting orders of yesterday")
		return
	}
	if len(ordersOfYesterday) > 0 {
		log.Info().Err(err).Int("item_id", itemID).Str("item_name", itemName).Msg("The item has been queried and stored")
		return
	}

	token, encryptedItemID, err := uc.Repo.GetTokenAndEncryptedItemID(dto.PORING_SERVER, 7, itemID, itemName)
	if err != nil {
		log.Error().Err(err).Int("item_id", itemID).Str("item_name", itemName).Msg("Error getting token")
		return
	}
	log.Info().Int("item_id", itemID).Str("item_name", itemName).Msgf("token: %s, encryptedItemID: %v\n", *token, *encryptedItemID)

	transactionCount, err := uc.Repo.GetTransactionCount(dto.PORING_SERVER, *encryptedItemID)
	if err != nil {
		log.Error().Err(err).Int("item_id", itemID).Str("item_name", itemName).Msg("Error getting transaction count")
		return
	}
	log.Info().Int("item_id", itemID).Str("item_name", itemName).Msgf("transaction count: %v\n", *transactionCount)

	transactions := make([]dto.TransactionsWithinIntervalEntry, 0)
	for start := 1; start <= *transactionCount; start += 30 {
		log.Info().Int("item_id", itemID).Str("item_name", itemName).Msgf("crawl transaction: %03d/%v\n", start, *transactionCount)
		pTransactions, err := uc.Repo.GetAllTransactionsWithinInterval(dto.PORING_SERVER, *encryptedItemID, *token, start)
		if err != nil {
			log.Error().Int("item_id", itemID).Str("item_name", itemName).Err(err).Msgf("Error getting transactions %03d/%v\n", start, *transactionCount)
			return
		}
		for i := range pTransactions {
			pTransaction := pTransactions[i]
			if pTransaction.Date == date {
				transactions = append(transactions, pTransaction)
			}
		}
	}
	// log.Info().Int("item_id", itemID).Str("item_name", itemName).Msgf("transactions: %+v\n", transactions)

	for _, t := range transactions {
		if err := uc.Repo.InsertOrders(ctx, genrepo.InsertOrdersParams{
			Date:   date,
			ItemID: int32(itemID),
			Count:  int32(t.Count),
			Price:  int32(t.Price),
		}); err != nil {
			log.Error().Err(err).Int("item_id", itemID).Str("item_name", itemName).Interface("transaction", t).Msg("Error inserting order")
			return
		}
	}
	log.Info().Int("item_id", itemID).Str("item_name", itemName).Msg("done inserting transactions")
}
