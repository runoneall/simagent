package kvstore

import (
	"encoding/json"
	"os"
)

func save() error {
	f, err := os.OpenFile(kvfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	return json.NewEncoder(f).Encode(store)
}
