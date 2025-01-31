package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"Groupie_Tracker/GetAPI"
)

type ArtistData struct {
	Artist    GetAPI.ArtistAPI
	Locations GetAPI.LocationsAPI
	Dates     GetAPI.DatesAPI
	Relations GetAPI.RelationAPI
}

type Server struct {
	Artists       []GetAPI.ArtistAPI
	ArtistDataMap map[int]ArtistData
	ArtistsToShow int
}

func NewServer(artistsToShow int) *Server {
	return &Server{
		ArtistsToShow: artistsToShow,
		ArtistDataMap: make(map[int]ArtistData),
	}
}

func (s *Server) LoadData(client *GetAPI.APIClient) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 4)

	wg.Add(1)

	go func() {
		defer wg.Done()
		var artists []GetAPI.ArtistAPI
		if err := client.Fetch("/api/artists", &artists); err != nil {
			errChan <- fmt.Errorf("artists fetch error: %v", err)
			return
		}
		s.Artists = artists

		for _, artist := range s.Artists {
			var artistData ArtistData
			artistData.Artist = artist

			locationPath := artist.Locations
			if len(locationPath) > 0 && locationPath[:4] == "http" {
				locationPath = artist.Locations[len("https://groupietrackers.herokuapp.com"):]
			}

			if err := client.Fetch(locationPath, &artistData.Locations); err != nil {
				errChan <- fmt.Errorf("locations fetch error: %v", err)
				return
			}

			datesPath := artist.ConcertDates
			if len(datesPath) > 0 && datesPath[:4] == "http" {
				datesPath = artist.ConcertDates[len("https://groupietrackers.herokuapp.com"):]
			}

			if err := client.Fetch(datesPath, &artistData.Dates); err != nil {
				errChan <- fmt.Errorf("dates fetch error: %v", err)
				return
			}

			relationsPath := artist.Relations
			if len(relationsPath) > 0 && relationsPath[:4] == "http" {
				relationsPath = artist.Relations[len("https://groupietrackers.herokuapp.com"):]
			}

			if err := client.Fetch(relationsPath, &artistData.Relations); err != nil {
				errChan <- fmt.Errorf("relations fetch error: %v", err)
				return
			}

			s.ArtistDataMap[artist.ID] = artistData
		}
	}()

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("multiple errors occurred: %v", errors)
	}

	return nil
}

func (s *Server) StartServer() {
	funcMap := template.FuncMap{
		"removeAsterisks": removeAsterisks,
		"formatDate":      formatDate,
		"formatLocation":  formatLocation,
	}

	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	artistTmpl, err := template.New("artistPage.gohtml").Funcs(funcMap).ParseFiles("templates/artistPage.gohtml")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(s.Artists), func(i, j int) { s.Artists[i], s.Artists[j] = s.Artists[j], s.Artists[i] })
		artistsToShow := s.Artists
		if len(s.Artists) > s.ArtistsToShow {
			artistsToShow = s.Artists[:s.ArtistsToShow]
		}
		if err := tmpl.Execute(w, struct{ Artists []GetAPI.ArtistAPI }{Artists: artistsToShow}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.Artists)
	})

	http.HandleFunc("/api/locations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var locations []GetAPI.LocationsAPI
		for _, data := range s.ArtistDataMap {
			locations = append(locations, data.Locations)
		}
		json.NewEncoder(w).Encode(locations)
	})

	http.HandleFunc("/api/dates", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dates []GetAPI.DatesAPI
		for _, data := range s.ArtistDataMap {
			dates = append(dates, data.Dates)
		}
		json.NewEncoder(w).Encode(dates)
	})

	http.HandleFunc("/api/relations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var relations []GetAPI.RelationAPI
		for _, data := range s.ArtistDataMap {
			relations = append(relations, data.Relations)
		}
		json.NewEncoder(w).Encode(relations)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		apiURLs := map[string]string{
			"artists":   "https://groupietrackers.herokuapp.com/api/artists",
			"locations": "https://groupietrackers.herokuapp.com/api/locations",
			"dates":     "https://groupietrackers.herokuapp.com/api/dates",
			"relation":  "https://groupietrackers.herokuapp.com/api/relation",
		}
		json.NewEncoder(w).Encode(apiURLs)
	})

	http.HandleFunc("/homePage", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, s); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/artist/"):]
		artistID, err := strconv.Atoi(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		artistData, ok := s.ArtistDataMap[artistID]
		if !ok {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		if err := artistTmpl.Execute(w, artistData); err != nil {
			log.Printf("template execution error: %v", err)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("templates/about.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, s); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/filter", s.FilterHandler)

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func removeAsterisks(s string) string {
	return strings.Replace(s, "*", "", -1)
}

func (s *Server) FilterHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	creationDate := r.URL.Query().Get("creation_date")
	firstAlbumDate := r.URL.Query().Get("first_album_date")
	members := r.URL.Query()["members"]
	location := r.URL.Query().Get("location")

	var filteredArtists []GetAPI.ArtistAPI

	for _, artist := range s.Artists {
		if search != "" && !strings.Contains(strings.ToLower(artist.Name), strings.ToLower(search)) {
			continue
		}
		if creationDate != "" {
			year, err := strconv.Atoi(creationDate)
			if err != nil || artist.CreationDate != year {
				continue
			}
		}
		if firstAlbumDate != "" && !strings.Contains(artist.FirstAlbum, firstAlbumDate) {
			continue
		}
		if len(members) > 0 {
			memberCount := strconv.Itoa(len(artist.Members))
			if !contains(members, memberCount) {
				continue
			}
		}
		if location != "" && !strings.Contains(strings.ToLower(artist.Locations), strings.ToLower(location)) {
			continue
		}
		filteredArtists = append(filteredArtists, artist)
	}

	tmpl, err := template.ParseFiles("templates/filter.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, struct{ Artists []GetAPI.ArtistAPI }{Artists: filteredArtists}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func formatDate(date string) string {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return date
	}

	monthMap := map[string]string{
		"01": "January", "02": "February", "03": "March", "04": "April",
		"05": "May", "06": "June", "07": "July", "08": "August",
		"09": "September", "10": "October", "11": "November", "12": "December",
	}

	day := parts[0]
	month := monthMap[parts[1]]
	year := parts[2]

	return fmt.Sprintf("%s %s %s", day, month, year)
}

func formatLocation(location string) string {
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.ReplaceAll(location, "-", " ")

	words := strings.Fields(location)
	if len(words) < 2 {
		return strings.Title(location)
	}

	for i, word := range words {
		words[i] = strings.Title(word)
	}

	city := strings.Join(words[:len(words)-1], " ")
	countryOrContinent := strings.ToUpper(words[len(words)-1])

	return fmt.Sprintf("%s, %s", city, countryOrContinent)
}
