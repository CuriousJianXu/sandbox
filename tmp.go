package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func tmp() {
	// The URL to which you want to send the POST request
	url := "https://event.gnjoy.com.tw/Ro/RoShopSearch/forAjax_shopDeal"

	// The payload you want to send in the request (in JSON format, for example)
	payload := []byte(`{"div_svr":"4290","div_storetype":"2","txb_KeyWord":"強烈","row_start":"1","recaptcha":"","sort_by":"itemPrice","sort_desc":""}`)

	// Create a request with the HTTP method "POST" and the request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the "Content-Type" header to specify the data format you are sending
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.Status == "200 OK" {
		fmt.Println("Request was successful")
	} else {
		fmt.Println("Request failed with status:", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	fmt.Printf("Request was successful %s", string(bodyBytes))

	// You can now read and process the response from the server if needed
}
