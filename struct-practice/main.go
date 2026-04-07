package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/struct-practice/note"
	todo "example.com/struct-practice/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	Display()
	saver
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text : ")

	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Failed to create todo:", err)
		return
	}
	err = outputItem(todoItem)
	if err != nil {
		fmt.Println("Failed to save todo:", err)
		return
	}

	noteItem, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}
	err = outputItem(noteItem)
	if err != nil {
		fmt.Println("Failed to save note:", err)
		return
	}
}

func outputItem(data outputtable) error {
	data.Display()
	err := data.Save()
	if err != nil {
		fmt.Println("Failed to save item:", err)
		return err
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // Windows対応
	return text
}

func printSomeThing(value interface{}) {
	typedValue, ok := value.(string)
	if ok {
		fmt.Println("String value:", typedValue)
	} else {
		fmt.Println("Unknown type")
	}
	// switch value.(type) {
	// case string:
	// 	fmt.Println("String value:", value)
	// case int:
	// 	fmt.Println("Integer value:", value)
	// default:
	// 	fmt.Println("Unknown type")
	// }
}
