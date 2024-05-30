package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func GetPosts() ([]byte, error) {
// 	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
// 	// resp, err := http.Get("https://sample-go.free.beeceptor.com/posts")
// 	resp, err := http.Get("https://662e647da7dda1fa378cd378.mockapi.io/api/v1/go-test")
// 	if err!= nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err!= nil {
// 		return nil, err
// 	}

// 	return body, nil
// }

// func GetPost() ([]byte, error) {
// 	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
// 	// resp, err := http.Get("https://sample-go.free.beeceptor.com/posts")
// 	resp, err := http.Get("https://662e647da7dda1fa378cd378.mockapi.io/api/v1/go-test")
// 	if err!= nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err!= nil {
// 		return nil, err
// 	}

// 	return body, nil
// }

// func ReadPostHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Received request: %s\n", r.URL.Path)

// 	// Set CORS headers
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	getResp, err := GetPost()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to get post: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(getResp)
// }