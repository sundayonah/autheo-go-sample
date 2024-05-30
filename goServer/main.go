// 	// postToUpdate := PostData{
// 	// 	ID: 12,
//     //     Title: "update post",
//     //     Body: "To use the UpdatePost function, you would pass an instance of PostData that includes the ID of the post you want to update along with any other fields you wish to change. Here's an example of how you might call this function:",
//     //     UserID: 1,
// 	// }
// 	// err := UpdatePost(postToUpdate)
// 	// if err!= nil {
//     //     fmt.Println("Error updating resource:", err)
//     // } else {
// 	// 	fmt.Println("Resource updated successfully.")
// 	// }
// 	//////////////////////////////////////////////////////

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostData struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Body   string             `json:"body"`
}

var client *mongo.Client
var postCollection *mongo.Collection


func main() {
	// Set up MongoDB connection
	MONGODB_URI := "mongodb+srv://phindCode:phindCode@cluster0.kcfnncd.mongodb.net/vite-with-go"
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	postCollection = client.Database("testdb").Collection("posts")

	http.HandleFunc("/api/go-test", ReadPostHandler)
	http.HandleFunc("/api/create-post", CreatePostHandler)
	http.HandleFunc("/api/delete-post", DeletePostHandler)


	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s\n", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var post PostData
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode post: %v", err), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.InsertOne(ctx, post)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert post: %v", err), http.StatusInternalServerError)
		return
	}

	post.ID = result.InsertedID.(primitive.ObjectID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func ReadPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s\n", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := postCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch posts: %v", err), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var posts []PostData
	if err = cursor.All(ctx, &posts); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode posts: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s\n", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the ID from the query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete post: %v", err), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "No post found with the given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Post with ID %s deleted", id)))
}



func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s\n", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var post PostData
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode post: %v", err), http.StatusBadRequest)
		return
	}

	// Ensure the ID is provided
	if post.ID == primitive.NilObjectID {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": bson.M{"title": post.Title, "body": post.Body}}

	result, err := postCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update post: %v", err), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "No post found with the given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}