package GetAPI

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
	Spotify      string   `json:"spotifyUrl"`
}
type SpotifyAPI struct {
	id      int    `json:"id"`
	spotify string `json:"spotifyUrl"`
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
	ID             int      `json:"id"`
	DatesLocations DatesAPI `json:"datesLocations"`
}

type CombinedData struct {
	Artist    ArtistAPI
	Locations LocationsAPI
	Dates     DatesAPI
	Relations RelationAPI
}
