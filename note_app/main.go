package main

import (
	"fmt"
	"note-app/note"
)

func main() {
	title, content := getUserNote()
	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.ToString()
}

func getUserNote() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
