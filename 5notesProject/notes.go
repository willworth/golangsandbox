package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/example/5notesProject/note"
	"example.com/example/5notesProject/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("err")
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("err")
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}
	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the file failed")
		return err
	}
	fmt.Println("saving the file succeeded")
	return nil
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
