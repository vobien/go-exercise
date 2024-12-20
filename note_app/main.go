package main

import (
	"bufio"
	"fmt"
	"note-app/note"
	"os"
	"strings"
)

func main() {
	title, content := getUserNote()
	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.ToString()
	note.Save()
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
