package GetAPI

type ArtistAPI struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationsAPI struct {
	Index []struct {
		ID   int      `json:"id"`
		Locs []string `json:"locations"`
	} `json:"index"`
}

type DatesAPI struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type RelationAPI struct {
	Index []struct {
		ID  int                 `json:"id"`
		Rel map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type CombinedData struct {
	Artist ArtistAPI
	Locs   []string
	Dates  []string
	Rel    map[string][]string
}
