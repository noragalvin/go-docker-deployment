package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}

	http.HandleFunc("/", hello)

	fmt.Println("Server started on port 8888")
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "Success",
		Message: "Hello world!",
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
