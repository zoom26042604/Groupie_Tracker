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
