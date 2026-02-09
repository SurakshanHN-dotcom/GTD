package internal

import (
	"encoding/json"
	"os"
)

const filePath = "data/todos.json"

type Store struct {
	NextID int    `json:"next_id"`
	Todos  []Todo `json:"todos"`
}

func LoadStore() (Store, error) {
	store := Store{
		NextID: 1,
		Todos:  []Todo{},
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return store, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return store, err
	}

	err = json.Unmarshal(data, &store)
	return store, err
}

func SaveStore(store Store) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	os.MkdirAll("data", 0755)
	return os.WriteFile(filePath, data, 0644)
}
