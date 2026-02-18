package service

import (
	"fmt"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/filesystem"
)

func AddChapter(name, filepath string) error {
	if err := filesystem.FindBookRoot(); err != nil {
		return err
	}

	var book book.Book
	book.UnmarshalBook()

	fmt.Println(book)

	return nil
}