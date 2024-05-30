package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func DeletePost(id int) error {
    url := "https://662e647da7dda1fa378cd378.mockapi.io/api/v1/go-test" + strconv.Itoa(id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err!= nil {
        return err
    }
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