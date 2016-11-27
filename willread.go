package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Link struct {
	Title string   `json:"title"`
	Link  string   `json:"link"`
	Tag   []string `json:"tag,string"`
}

func newLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var rec Link
	err := decoder.Decode(&rec)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Println(rec)
}

func main() {
	http.HandleFunc("/new", newLink)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
