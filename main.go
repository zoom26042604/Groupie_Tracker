package main

import (
	"Groupie_Tracker/GetAPI"
	"Groupie_Tracker/handler"
	"log"
)

func main() {
	server := handler.NewServer(12)
	apiClient := GetAPI.NewAPIClient("https://groupietrackers.herokuapp.com")

	if err := server.LoadData(apiClient); err != nil {
		log.Fatalf("Failed to fetch initial data: %v", err)
	}

	server.StartServer()
}
