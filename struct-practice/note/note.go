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
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Println("Title : ", note.Title)
	fmt.Println("Content : ", note.Content)
	fmt.Println("Created At : ", note.CreatedAt)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		fmt.Println("Failed to serialize note:", err)
		return err
	}
	return os.WriteFile(fileName, []byte(jsonData), 0644)
}
