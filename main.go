package main

import (
	"go-docker-development/app/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/user", controllers.Index)

	port := os.Getenv("PORT")
	log.Println("Server is running on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println(err)
	}

}
