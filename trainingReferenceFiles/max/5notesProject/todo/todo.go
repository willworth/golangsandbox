package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// receiver shows what struct the method is attached to
// receiver can be named, allowing use within method
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // returns error if fails

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
