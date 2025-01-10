package main

import (
	"log"
	"serveurTest/GetAPI"
	"serveurTest/handler"
)

func main() {
	server := handler.NewServer()
	apiClient := GetAPI.NewAPIClient("https://groupietrackers.herokuapp.com")

	if err := server.LoadData(apiClient); err != nil {
		log.Fatalf("Failed to fetch initial data: %v", err)
	}

	server.StartServer()
}
