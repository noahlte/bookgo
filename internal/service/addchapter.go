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

func AddChapter(newChapter *book.Chapter) error {
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

	foldername := filesystem.RenameFile(newChapter.Name)
	
	filepath := fmt.Sprintf("%d-chapter-%s", chapterNumber, foldername)

	if _, err := os.Stat(path.Join(util.ContentFile, filepath)); err == nil {
		return errors.New("this chapter already exist")
	}

	newChapter.Number = chapterNumber
	newChapter.Path = path.Join(bookPath, util.ContentFile, filepath)

	userBook.Chapters = append(userBook.Chapters, *newChapter)

	err = os.Mkdir(path.Join(util.ContentFile, filepath), 0755)
	if err != nil {
		return err
	}

	err = userBook.Save()
	
	return nil
}