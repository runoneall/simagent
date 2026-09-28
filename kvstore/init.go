package kvstore

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const kvfile = "kvstore.json"

var (
	store = map[string]string{}
	lock  sync.RWMutex
)

func init() {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.OpenFile(kvfile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	if err := json.NewDecoder(f).Decode(&store); err != nil {
		if err := f.Truncate(0); err != nil {
			log.Fatalln(err)
		}

		if _, err := f.Seek(0, 0); err != nil {
			log.Fatalln(err)
		}

		if err := json.NewEncoder(f).Encode(store); err != nil {
			log.Fatalln(err)
		}
	}
}
