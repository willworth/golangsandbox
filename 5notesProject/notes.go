package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/example/5notesProject/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("err")
		return
	}

	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	//specify the source - here it's the command line
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // where to stop

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") //WINDOWS

	return text

}
