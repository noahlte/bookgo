package service

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/filesystem"
	"github.com/noahlte/bookgo/internal/util"
)

func SetupBook(name, author string) error {
	filepath := filesystem.RenameFile(name)

	if _, err := os.Stat(filepath); err == nil {
		return errors.New("book files already exist")
	}

	userBook := &book.Book{
		Name: name,
		Description: "Description...",
		Author: author,
		CreatedAt: time.Now(),
	}

	err := os.Mkdir(filepath, 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(filepath, util.ContentFile), 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(filepath, util.ImagesFile), 0755)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(userBook)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(filepath, util.YamlFile), data, 0644)
	if err != nil {
		return err
	}
	
	fmt.Printf("Your book %s has been created !\n", userBook.Name)

	return nil
}