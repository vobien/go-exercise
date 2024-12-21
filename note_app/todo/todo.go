package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Text   string    `json:"text"`
	CreatedAt time.Time `json:"createAt"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("text must be non empty")
	}

	return Todo{
		Text:     text,
		CreatedAt: time.Now(),
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Todo: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}
