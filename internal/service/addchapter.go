package service

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/filesystem"
	"github.com/noahlte/bookgo/internal/util"
)

func AddChapter(name string) error {
	if err := filesystem.FindBookRoot(); err != nil {
		return err
	}

	var userBook book.Book
	userBook.UnmarshalBook()

	bookPath, err := os.Getwd()
	if err != nil {
		return err
	}

	chapterNumber := len(userBook.Chapters) + 1

	foldername := filesystem.RenameFile(name)
	
	filepath := fmt.Sprintf("%d-chapter-%s", chapterNumber, foldername)

	if _, err := os.Stat(path.Join(util.ContentFile, filepath)); err == nil {
		return errors.New("this chapter already exist")
	}

	chapter := &book.Chapter{
		Name: name,
		Description: "...",
		Number: chapterNumber,
		Path: path.Join(bookPath, util.ContentFile, filepath),
	}

	userBook.Chapters = append(userBook.Chapters, *chapter)

	err = os.Mkdir(path.Join(util.ContentFile, filepath), 0755)
	if err != nil {
		return err
	}

	err = userBook.Save()
	
	return nil
}