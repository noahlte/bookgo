package service

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/noahlte/bookgo/internal/book"
)

var restrictedChar = []string{" ", "?", "!", ".", "/", ":", ",", "€", "$", "+", "-", "=", "*", "µ", "¨", "^", "°", "'"} 

func SetupBook(name, author, filepath string) error {
	filepath = strings.ToLower(filepath)
	for _, char := range restrictedChar {
		filepath = strings.ReplaceAll(filepath, char, "-")
	}

	if _, err := os.Stat(filepath); err == nil {
		return errors.New("book files already exist")
	}

	book := &book.Book{
		Name: name,
		Description: "Description...",
		Author: author,
		CreatedAt: time.Now(),
	}

	err := os.Mkdir(filepath, 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(filepath, "content"), 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(filepath, "images"), 0755)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(book)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(filepath, "book.yaml"), data, 0644)
	if err != nil {
		return err
	}
	
	fmt.Printf("Your book %s has been created !", book.Name)

	return nil
}