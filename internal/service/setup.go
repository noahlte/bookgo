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

func SetupBook(newBook book.Book) error {
	filepath := filesystem.RenameFile(newBook.Name)

	if _, err := os.Stat(filepath); err == nil {
		return errors.New("book files already exist")
	}

	newBook.CreatedAt = time.Now()

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

	data, err := yaml.Marshal(newBook)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(filepath, util.YamlFile), data, 0644)
	if err != nil {
		return err
	}
	
	fmt.Printf("Your book %s has been created !\n", newBook.Name)

	return nil
}