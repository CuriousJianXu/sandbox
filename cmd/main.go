package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"oproaster.com/sandbox/dto"
)

func main() {
	token, encryptedItemID, err := getTokenAndEncryptedItemID(dto.PORING_SERVER, 7, 25375, "強烈靈魂精髓")
	if err != nil {
		fmt.Println("Error getting token:", err)
		return
	}
	// token_ := "020d949b2810"
	// token := &token_
	fmt.Printf("token: %s, encryptedItemID: %v\n", *token, *encryptedItemID)
	time.Sleep(1 * time.Second)
	transactionCount, err := getTransactionCount(dto.PORING_SERVER, *encryptedItemID)
	if err != nil {
		fmt.Println("Error getting transaction count:", err)
		return
	}
	fmt.Printf("transaction count: %v\n", *transactionCount)
	*transactionCount = *transactionCount / 2
}

func getTokenAndEncryptedItemID(server string, days int, itemID int, itemName string) (*string, *string, error) {
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

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
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
	if err := decoder.Decode(&data); err != nil {
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

// GetLocalTimeInTaipei returns the local time in "Asia/Taipei" timezone for a specified
// number of days before the current date, formatted as "yyyy/mm/dd"
func GetLocalTimeInTaipei(daysBefore int) (string, error) {
	// Set the timezone to "Asia/Taipei" (GMT+8)
	taipeiLocation, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return "", err
	}

	// Get the current local time in "Asia/Taipei" timezone
	localTime := time.Now().AddDate(0, 0, -daysBefore).In(taipeiLocation)

	// Format the local time as "yyyy/mm/dd"
	formattedTime := localTime.Format("2006/01/02")

	return formattedTime, nil
}

func getYesterdayTransactions(server string, token string, encryptedItemID string) (*string, *string, error) {
	// Get transaction counts
	// list transactions per day with pagination.
	return nil, nil, nil
}

func getTransactionCount(server string, encryptedItemID string) (*int, error) {
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

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
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
	cnt, err := strconv.Atoi(strings.Replace(string(bodyBytes), "\"", "", -1))
	if err != nil {
		return nil, fmt.Errorf("The response is not a integer: %v, error: %w", string(bodyBytes), err)
	}
	return &cnt, nil
}

// func getYesterdayTransactions(server string, token string, encryptedItemID string) (*string, *string, error) {
// 	url := "https://event.gnjoy.com.tw/Ro/RoShopSearch/forAjax_history_a_item_DealDetail"

// 	payload := dto.TransactionPerDayPayload{
// 		Server:          server,
// 		Days:            "7",
// 		EncryptedItemID: encryptedItemID,
// 		Token:           token,
// 	}

// 	payloadBytes, err := json.Marshal(payload)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Error marshaling JSON: %w", err)
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Error creating request: %w", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

// 	// Create an HTTP client
// 	client := &http.Client{}

// 	// Send the request
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Error sending request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	// Check the response status code
// 	if resp.StatusCode != http.StatusOK {
// 		bodyBytes, _ := io.ReadAll(resp.Body)
// 		return nil, nil, fmt.Errorf("Request failed with status: %v, message: %s", resp.StatusCode, string(bodyBytes))
// 	}

// 	// Decode the JSON response into your struct
// 	data := dto.HistoryResp{}
// 	decoder := json.NewDecoder(resp.Body)
// 	if err := decoder.Decode(&data); err != nil {
// 		return nil, nil, fmt.Errorf("Error sending request: %w", err)
// 	}

// 	encryptedItemID, err := findEntryByItemID(itemID, data.Data)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Error finding entry by item id: %w", err)
// 	}

// 	return &data.Token, encryptedItemID, nil
// }
