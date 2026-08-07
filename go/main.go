package main

import (
	"log"
	"os"

	"github.com/Yustinia/gopaper"
)

var (
	APIKEY = os.Getenv("WALLHAVEN_API_KEY")
	CLIENT = gopaper.NewClient(APIKEY)
)

func main() {
	params := gopaper.NewSearch()
	params.KeySearch = "japan"

	result, err := CLIENT.FetchPages(&params, 1, 2)
	if err != nil {
		log.Println(err)
	}

	for _, wall := range result {
		log.Println(wall.Path)
	}
}
