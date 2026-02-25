package service

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/util"
)

type readMeTemplate struct {
	BookName	string
	BookPath 	string
}

//go:embed templates/*
var readmeTemplate embed.FS

func SetupBook(newBook *book.Book) error {
	bookPath := util.SanitizeName(newBook.Name)

	if _, err := os.Stat(bookPath); err == nil {
		return errors.New("book files already exist")
	}

	newBook.CreatedAt = time.Now()

	err := os.Mkdir(bookPath, 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(filepath.Join(bookPath, util.ContentDir), 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(filepath.Join(bookPath, util.ImagesDir), 0755)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(newBook)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(bookPath, util.YamlFile), data, 0644)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFS(readmeTemplate, "templates/README.md")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.Join(bookPath, "README.md"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, readMeTemplate{BookName: newBook.Name, BookPath: bookPath})
	if err != nil {
		return err
	}

	fmt.Printf("Your book %s has been created !\n", newBook.Name)

	return nil
}
