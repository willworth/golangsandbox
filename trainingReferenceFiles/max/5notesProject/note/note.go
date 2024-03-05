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
	Title     string    `json:"title"` //struct tags
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// receiver shows what struct the method is attached to
// receiver can be named, allowing use within method
func (nameHere Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v", nameHere.Title, nameHere.Content)
}

func (internalNote Note) Save() error {
	fileName := strings.ReplaceAll(internalNote.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(internalNote) // only if contents is made public (ie upper case)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // returns error if fails

}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
