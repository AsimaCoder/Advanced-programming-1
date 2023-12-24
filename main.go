package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Установка заголовка CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		response := Response{"400", "Invalid JSON message"}
		w.WriteHeader(http.StatusBadRequest) // Установка статуса ответа
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Printf("Message received: %s\n", req.Message)

	response := Response{"success", "Data successfully received"}
	json.NewEncoder(w).Encode(response)
}
