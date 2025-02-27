package GetAPI

import (
	"fmt"
	"strings"
)

type ArtistAPI struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	SpotifyURL   string   `json:"spotifyUrl"`
	SpotifyID    string   `json:"spotifyId"`
}

type LocationsAPI struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesAPI struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationAPI struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistData struct {
	ID         string `json:"id"`
	SpotifyURL string `json:"spotifyUrl"`
	MusicUrl   string `json:"musicUrl"`
}

func ShowArtistData(artist ArtistAPI) {
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("Artist ID: %d\n", artist.ID)
	fmt.Printf("Artist Image: %s\n", artist.Image)
	fmt.Printf("Artist Name: %s\n", artist.Name)
	fmt.Printf("Artist Members: %s\n", strings.Join(artist.Members, ","))
	fmt.Printf("Artist Creation Date: %s\n", artist.CreationDate)
	fmt.Printf("Artist First Album: %s\n", artist.FirstAlbum)
	fmt.Printf("Artist Locations: %s\n", artist.Locations)
	fmt.Printf("Artist Concert Dates: %s\n", artist.ConcertDates)
	fmt.Printf("Artist Relations: %s\n", artist.Relations)
	fmt.Printf("Artist Spotify URL: %s\n", artist.SpotifyURL)
	fmt.Printf("Artist Spotify ID: %s\n", artist.SpotifyID)
	fmt.Println("----------------------------------------------------------")
}

type CombinedData struct {
	Artist    ArtistAPI
	Locations LocationsAPI
	Dates     DatesAPI
	Relations RelationAPI
}
