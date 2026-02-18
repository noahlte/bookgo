package service

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/noahlte/bookgo/internal/book"
)

// TODO: Setup command service
func SetupBook(name, author string) error {
	book := &book.Book{
		Name: name,
		Description: "Description...",
		Author: author,
		CreatedAt: time.Now(),
	}

	fmt.Printf(
		"Creating a new book : %s writed by %s. This book talk about %s and was created %v\n", 
		book.Name, 
		book.Author, 
		book.Description, 
		book.CreatedAt,
	)

	data, err := yaml.Marshal(book)
	if err != nil {
		return err
	}

	err = os.WriteFile("book.yaml", data, 6444)
	if err != nil {
		return err
	}

	return nil
}