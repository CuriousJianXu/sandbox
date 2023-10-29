package repo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"oproaster.com/sandbox/dto"
	"oproaster.com/sandbox/utils"
)

func (repo *Repo) GetTokenAndEncryptedItemID(server string, days int, itemID int, itemName string) (*string, *string, error) {
	utils.FriendlyCrawl()

	url := "https://event.gnjoy.com.tw/Ro/RoShopSearch/forAjax_history"

	payload := dto.HistoryPayload{
		Server:    server,
		Days:      strconv.Itoa(days),
		Keyword:   itemName,
		Recaptcha: "",
		SortBy:    "SumitemCNT",
		SortDesc:  "desc",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, fmt.Errorf("Error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, nil, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("Error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, nil, fmt.Errorf("Request failed with status: %v, message: %s", resp.StatusCode, string(bodyBytes))
	}

	// Decode the JSON response into your struct
	data := dto.HistoryResp{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&data); err != nil {
		return nil, nil, fmt.Errorf("Error sending request: %w", err)
	}

	encryptedItemID, err := findEntryByItemID(itemID, data.Data)
	if err != nil {
		return nil, nil, fmt.Errorf("Error finding entry by item id: %w", err)
	}

	return &data.Token, encryptedItemID, nil
}

func findEntryByItemID(itemID int, entryList dto.HistoryData) (*string, error) {
	for _, entry := range entryList {
		if entry.ItemID == itemID {
			return &entry.EncryptedItemID, nil
		}
	}
	return nil, errors.New("itemID not found")
}

func (repo *Repo) GetTransactionCount(server string, encryptedItemID string) (*int, error) {
	utils.FriendlyCrawl()
	url := "https://event.gnjoy.com.tw/Ro/RoShopSearch/forAjax_history_a_item_DealDetail_count"

	payload := dto.TransactionCountPayload{
		Server:          server,
		Days:            "7",
		EncryptedItemID: encryptedItemID,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Request failed with status: %v, message: %s", resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	// The response is in this format: "126"
	// So we need to remove the " in the response the get transaction count
	cnt, err := strconv.Atoi(strings.ReplaceAll(string(bodyBytes), "\"", ""))
	if err != nil {
		return nil, fmt.Errorf("The response is not a integer: %v, error: %w", string(bodyBytes), err)
	}
	return &cnt, nil
}

func (repo *Repo) GetAllTransactionsWithinInterval(server string, encryptedItemID string, token string, start int) ([]dto.TransactionsWithinIntervalEntry, error) {
	utils.FriendlyCrawl()

	url := "https://event.gnjoy.com.tw/Ro/RoShopSearch/forAjax_history_a_item_DealDetail"

	payload := dto.TransactionsWithinIntervalPayload{
		Server:          server,
		Days:            "7",
		RowStart:        strconv.Itoa(start),
		EncryptedItemID: encryptedItemID,
		Token:           token,
		SortBy:          "regDate_",
		SortDesc:        "desc",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Request failed with status: %v, message: %s", resp.StatusCode, string(bodyBytes))
	}

	// Decode the JSON response into your struct
	data := dto.TransactionsWithinIntervalResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	return data.Data, nil
}
