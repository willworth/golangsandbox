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

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

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

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

// func printSomething(value interface{}) {  // is the same as:
func printSomething(value any) {

	intVal, ok := value.(int)
	if ok {
		fmt.Println("integer", intVal)
		return
	}

	float64Val, ok := value.(float64)
	if ok {
		fmt.Println("float64", float64Val)
		return

	}

	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	}

	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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