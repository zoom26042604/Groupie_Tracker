package GetAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *APIClient) Fetch(endpoint string, target interface{}) error {
	url := c.BaseURL + endpoint
	fmt.Printf("Fetching from: %s\n", url)

	resp, err := c.Client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("API returned status code %d for %s. Response: %s", resp.StatusCode, url, string(bodyBytes))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetCombinedData(artists []ArtistAPI, locations LocationsAPI, dates DatesAPI, relations RelationAPI) []CombinedData {
	var combinedData []CombinedData
	for _, artist := range artists {
		combinedData = append(combinedData, CombinedData{
			Artist:    artist,
			Locations: locations,
			Dates:     dates,
			Relations: relations,
		})
	}
	return combinedData
}

func GetCombinedDataByID(artists []ArtistAPI, locations []LocationsAPI, dates []DatesAPI, relations []RelationAPI, id int) (CombinedData, error) {
	var combinedData CombinedData
	for _, artist := range artists {
		if artist.ID == id {
			combinedData.Artist = artist
			break
		}
	}
	if combinedData.Artist.ID == 0 {
		return combinedData, fmt.Errorf("artist with ID %d not found", id)
	}

	for _, location := range locations {
		if location.ID == id {
			combinedData.Locations = location
			break
		}
	}

	for _, date := range dates {
		if date.ID == id {
			combinedData.Dates = date
			break
		}
	}

	for _, relation := range relations {
		if relation.ID == id {
			combinedData.Relations = relation
			break
		}
	}

	return combinedData, nil
}
