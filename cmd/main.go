package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	fmt.Printf("token: %s, encryptedItemID: %v", *token, *encryptedItemID)
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
