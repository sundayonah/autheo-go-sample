package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// type PostData struct {
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// 	UserID int    `json:"userId"`
// }

func UpdatePost(data PostData) error {
	jsonData, err := json.Marshal(data)
	if err!= nil {
		return err
	}

	// Construct the URL with the post ID
	url := fmt.Sprintf("https://662e647da7dda1fa378cd378.mockapi.io/api/v1/go-test/%d", data.ID)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err!= nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err!= nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode!= http.StatusOK {
		return fmt.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	return nil
}
