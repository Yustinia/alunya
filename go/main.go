package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Yustinia/gopaper"
)

var (
	APIKEY = os.Getenv("WALLHAVEN_API_KEY")
	CLIENT = gopaper.NewClient(APIKEY)
)

type SearchFilters struct {
	Query      string `json:"query"`
	Categories string `json:"categories"`
	Purity     string `json:"purity"`
	Sorting    string `json:"sorting"`
	Order      string `json:"order"`
	TopRange   string `json:"toprange"`
	AtLeast    string `json:"atleast"`
	Resolution string `json:"resolution"`
	Ratios     string `json:"ratios"`
	Page       int    `json:"page"`
	Seed       string `json:"seed"`
}

func doSearch(w http.ResponseWriter, r *http.Request) {
	sf := SearchFilters{}
	sp := gopaper.NewSearch()

	err := json.NewDecoder(r.Body).Decode(&sf)
	if err != nil {
		http.Error(w, "failed to decode", http.StatusBadRequest)
		return
	}

	sp.KeySearch = sf.Query
	sp.Categories = sf.Categories
	sp.Purity = sf.Purity
	sp.Sorting = sf.Sorting
	sp.Order = sf.Order
	sp.TopRange = sf.TopRange
	sp.AtLeast = sf.AtLeast
	sp.Resolution = sf.Resolution
	sp.Ratios = sf.Ratios
	sp.Page = sf.Page
	sp.Seed = sf.Seed

	result, err := CLIENT.Search(sp)
	if err != nil {
		http.Error(w, "request failed", http.StatusBadGateway)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/search", doSearch)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
