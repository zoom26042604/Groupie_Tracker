package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	"serveurTest/GetAPI"
)

type Server struct {
	Artists   []GetAPI.ArtistAPI
	Locations GetAPI.LocationsAPI
	Dates     GetAPI.DatesAPI
	Relations GetAPI.RelationAPI
	Mu        sync.RWMutex
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) LoadData(client *GetAPI.APIClient) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 4)

	wg.Add(4)

	go func() {
		defer wg.Done()
		var artists []GetAPI.ArtistAPI
		if err := client.Fetch("/api/artists", &artists); err != nil {
			errChan <- fmt.Errorf("artists fetch error: %v", err)
			return
		}
		s.Mu.Lock()
		s.Artists = artists
		s.Mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		var locations GetAPI.LocationsAPI
		if err := client.Fetch("/api/locations", &locations); err != nil {
			errChan <- fmt.Errorf("locations fetch error: %v", err)
			return
		}
		s.Mu.Lock()
		s.Locations = locations
		s.Mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		var dates GetAPI.DatesAPI
		if err := client.Fetch("/api/dates", &dates); err != nil {
			errChan <- fmt.Errorf("dates fetch error: %v", err)
			return
		}
		s.Mu.Lock()
		s.Dates = dates
		s.Mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		var relations GetAPI.RelationAPI
		if err := client.Fetch("/api/relation", &relations); err != nil {
			errChan <- fmt.Errorf("relations fetch error: %v", err)
			return
		}
		s.Mu.Lock()
		s.Relations = relations
		s.Mu.Unlock()
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
	tmpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		s.Mu.RLock()
		defer s.Mu.RUnlock()
		if err := tmpl.Execute(w, s); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.Mu.RLock()
		json.NewEncoder(w).Encode(s.Artists)
		s.Mu.RUnlock()
	})

	http.HandleFunc("/api/locations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.Mu.RLock()
		json.NewEncoder(w).Encode(s.Locations)
		s.Mu.RUnlock()
	})

	http.HandleFunc("/api/dates", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.Mu.RLock()
		json.NewEncoder(w).Encode(s.Dates)
		s.Mu.RUnlock()
	})

	http.HandleFunc("/api/relations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.Mu.RLock()
		json.NewEncoder(w).Encode(s.Relations)
		s.Mu.RUnlock()
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
		s.Mu.RLock()
		defer s.Mu.RUnlock()
		if err := tmpl.Execute(w, s); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/artist/"):]
		s.Mu.RLock()
		defer s.Mu.RUnlock()
		var artist GetAPI.ArtistAPI
		for _, a := range s.Artists {
			if fmt.Sprintf("%d", a.ID) == id {
				artist = a
				break
			}
		}
		if artist.ID == 0 {
			http.NotFound(w, r)
			return
		}
		tmpl, err := template.ParseFiles("templates/artistPage.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, struct{ Artist GetAPI.ArtistAPI }{Artist: artist}); err != nil {
			log.Printf("template execution error: %v", err)
		}
	})

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
