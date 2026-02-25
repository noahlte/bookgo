package service

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/filesystem"
	"github.com/noahlte/bookgo/internal/util"
)

type sectionTemplate struct {
	ChapterName        string
	ChapterNumber 		 int
}

//go:embed templates/*
var templateFiles embed.FS

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

	chapterPath := fmt.Sprintf("%d-chapter-%s", chapterNumber, foldername)

	if _, err := os.Stat(filepath.Join(util.ContentDir, chapterPath)); err == nil {
		return errors.New("this chapter already exist")
	}

	newChapter.Number = chapterNumber
	newChapter.Path = filepath.Join(bookPath, util.ContentDir, chapterPath)

	userBook.Chapters = append(userBook.Chapters, *newChapter)

	err = os.Mkdir(filepath.Join(util.ContentDir, chapterPath), 0755)
	if err != nil {
		return err
	}

	newSection := sectionTemplate{
		ChapterName:        newChapter.Name,
		ChapterNumber: 			newChapter.Number,
	}

	tmpl, err := template.ParseFS(templateFiles, "templates/new-section.md")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.Join(util.ContentDir, chapterPath, "new-section.md"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, newSection)
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
