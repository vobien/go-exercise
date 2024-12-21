package main

import (
	"bufio"
	"fmt"
	"note-app/note"
	"note-app/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func main() {
	title, content := getUserNote()
	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)

	todoText := getUserInput("Input todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
	}

	outputData(todo)
}

func getUserNote() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) {
	err := data.Save()
	if err != nil {
		fmt.Printf("Error when saving data to file, %s\n", err)
		return
	}
	fmt.Println("Save to file successfully")
}

func outputData(data outputable) {
	data.Display()
	saveData(data)
}