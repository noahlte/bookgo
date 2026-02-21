package service

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/filesystem"
	"github.com/noahlte/bookgo/internal/util"
)

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

	newSection := &book.SectionTemplate{
		ChapterName:        newChapter.Name,
		ChapterDescription: newChapter.Description,
	}

	tmpl, err := template.ParseFS(templateFiles, "templates/new-section.md")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(util.ContentDir, filepath, "new-section.md"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, *newSection)
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
