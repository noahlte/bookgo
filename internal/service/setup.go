package service

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path"
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
	filepath := util.SanitizeName(newBook.Name)

	if _, err := os.Stat(filepath); err == nil {
		return errors.New("book files already exist")
	}

	newBook.CreatedAt = time.Now()

	err := os.Mkdir(filepath, 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(filepath, util.ContentDir), 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(filepath, util.ImagesDir), 0755)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(newBook)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(filepath, util.YamlFile), data, 0644)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFS(readmeTemplate, "templates/README.md")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(filepath, "README.md"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, readMeTemplate{BookName: newBook.Name, BookPath: filepath})
	if err != nil {
		return err
	}

	fmt.Printf("Your book %s has been created !\n", newBook.Name)

	return nil
}
