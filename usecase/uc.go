package usecase

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"oproaster.com/sandbox/dto"
	"oproaster.com/sandbox/repo"
	genrepo "oproaster.com/sandbox/repo/gen"
	"oproaster.com/sandbox/utils"
)

type Usecase struct {
	Repo *repo.Repo
}

func New(db *sqlx.DB) *Usecase {
	return &Usecase{
		Repo: repo.New(db),
	}
}

func (uc *Usecase) CrawlAndStoreTransactions() {
	itemID := 25375
	token, encryptedItemID, err := uc.Repo.GetTokenAndEncryptedItemID(dto.PORING_SERVER, 7, itemID, "強烈靈魂精髓")
	if err != nil {
		fmt.Println("Error getting token:", err)
		return
	}
	// token_ := "ae3b79c7cafd"
	// token := &token_
	// encryptedItemID_ := "zbhvGu8RiV0="
	// encryptedItemID := &encryptedItemID_
	fmt.Printf("token: %s, encryptedItemID: %v\n", *token, *encryptedItemID)

	transactionCount, err := uc.Repo.GetTransactionCount(dto.PORING_SERVER, *encryptedItemID)
	if err != nil {
		fmt.Println("Error getting transaction count:", err)
		return
	}
	fmt.Printf("transaction count: %v\n", *transactionCount)

	localDateYesterday, err := utils.GetLocalTimeInTaipei(1)
	if err != nil {
		fmt.Println("Error getting local date yesterday:", err)
		return
	}

	transactions := make([]dto.TransactionsWithinIntervalEntry, 0)
	for start := 1; start <= *transactionCount; start += 30 {
		fmt.Printf("crawl transaction: %03d/%v\n", start, *transactionCount)
		pTransactions, err := uc.Repo.GetAllTransactionsWithinInterval(dto.PORING_SERVER, *encryptedItemID, *token, start)
		if err != nil {
			fmt.Printf("Error getting transactions %03d/%v: %v\n", start, *transactionCount, err)
			return
		}
		for i := range pTransactions {
			pTransaction := pTransactions[i]
			if pTransaction.Date == localDateYesterday {
				transactions = append(transactions, pTransaction)
			}
		}
	}
	fmt.Printf("transactions: %+v\n", transactions)

	for _, t := range transactions {
		if err := uc.Repo.InsertOrders(context.Background(), genrepo.InsertOrdersParams{
			Date:   localDateYesterday,
			ItemID: int32(itemID),
			Count:  int32(t.Count),
			Price:  int32(t.Price),
		}); err != nil {
			fmt.Printf("Error inserting order: %v, order: %+v\n", err, t)
			return
		}
	}
	fmt.Printf("done inserting transactions")
}
