package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// receiver shows what struct the method is attached to
// receiver can be named, allowing use within method
func (nameHere Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v", nameHere.title, nameHere.content)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
