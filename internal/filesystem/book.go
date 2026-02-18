package filesystem

import (
	"errors"
	"os"
)

func FindBookRoot() error {
	if _, err := os.Stat("book.yaml"); err != nil {
		return errors.New("you need to be in the book file to run this command (cd your-book-name)")
	}

	return nil
}