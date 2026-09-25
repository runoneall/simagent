package config

import (
	"encoding/json"
	"log"
	"os"
)

var cfg Config

func init() {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		log.Fatalln(err)
	}
}
