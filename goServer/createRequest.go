package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func CreatePost(data PostData) error {
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	// req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com", bytes.NewBuffer(jsonData))
// 	req, err := http.NewRequest("POST", "https://662e647da7dda1fa378cd378.mockapi.io/api/v1/go-test", bytes.NewBuffer(jsonData))
//     if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err!= nil {
//         return err
//     }
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusCreated {
// 		return fmt.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
// 	}
// 	return nil
// }

// func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Received request: %s\n", r.URL.Path)

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var post PostData
// 	err := json.NewDecoder(r.Body).Decode(&post)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to decode post: %v", err), http.StatusBadRequest)
// 		return
// 	}

// 	// Forward the POST request to the mock API
// 	mockAPIURL := "https://662e647da7dda1fa378cd378.mockapi.io/api/v1/go-test"
// 	postBody, err := json.Marshal(post)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to marshal post: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	log.Printf("Received request: %s\n", postBody)

// 	resp, err := http.Post(mockAPIURL, "application/json", bytes.NewBuffer(postBody))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to forward post to mock API: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to read mock API response: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(resp.StatusCode)
// 	w.Write(body)
// }