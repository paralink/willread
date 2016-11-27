package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var store LinkStore

type Link struct {
	Title string   `json:"title"`
	Link  string   `json:"link"`
	Tag   []string `json:"tag"`
}

type LinkStore struct {
	links []*Link
}

func (s *LinkStore) Add(l *Link) {
	s.links = append(s.links, l)
}

func (s *LinkStore) List() []*Link {
	return s.links
}

func newLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var rec *Link

	if err := decoder.Decode(&rec); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	store.Add(rec)

	fmt.Fprintln(w, "Succeed")
}

func batchLinks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var links []*Link
	if err := decoder.Decode(&links); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for _, link := range links {
		store.Add(link)
	}

	fmt.Fprintln(w, "Succeed")
}

func listLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	links := store.List()
	json.NewEncoder(w).Encode(links)
}

func main() {
	http.HandleFunc("/new", newLink)
	http.HandleFunc("/list", listLink)
	http.HandleFunc("/new_batch", batchLinks)

	// storage
	log.Fatal(http.ListenAndServe(":8080", nil))
}
