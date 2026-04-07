package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content cannot be empty")
	}
	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println("Text : ", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Failed to serialize todo:", err)
		return err
	}
	return os.WriteFile(fileName, []byte(jsonData), 0644)
}
