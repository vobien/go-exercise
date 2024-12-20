package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createAt"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content must be non empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) ToString() {
	fmt.Printf("Title: %s, Content: %s\n", note.Title, note.Content)
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	filename = strings.ToLower(filename)

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}
