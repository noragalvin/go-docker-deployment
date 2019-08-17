package controllers

import (
	"encoding/json"
	"go-docker-development/app/models"
	"net/http"
)

// Index ..
func Index(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID:    1,
		Name:  "Bui Ngoc Minh",
		Email: "minhnora98@gmail.com",
	}
	status := true
	message := "success"
	data := map[string]interface{}{"status": status, "message": message}
	data["user"] = user
	json.NewEncoder(w).Encode(data)
}
