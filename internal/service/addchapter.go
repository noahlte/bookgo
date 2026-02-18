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
	err := userBook.UnmarshalBook()
	if err != nil {
		return err
	}

	bookPath, err := os.Getwd()
	if err != nil {
		return err
	}

	chapterNumber := len(userBook.Chapters) + 1

	foldername := util.SanitizeName(newChapter.Name)
	
	filepath := fmt.Sprintf("%d-chapter-%s", chapterNumber, foldername)

	if _, err := os.Stat(path.Join(util.ContentDir, filepath)); err == nil {
		return errors.New("this chapter already exist")
	}

	newChapter.Number = chapterNumber
	newChapter.Path = path.Join(bookPath, util.ContentDir, filepath)

	userBook.Chapters = append(userBook.Chapters, *newChapter)

	err = os.Mkdir(path.Join(util.ContentDir, filepath), 0755)
	if err != nil {
		return err
	}

	err = userBook.Save()
	if err != nil {
		return err
	}

	fmt.Printf("The chapter %s has been created!", newChapter.Name)
	
	return nil
}