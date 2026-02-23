package filesystem

import (
	"errors"
	"os"

	"github.com/noahlte/bookgo/internal/util"
)

// Chercher si le book.yaml existe, si non, cela veut dire que nous ne sommes pas dans un fichier livre.
func FindBookRoot() error {
	if _, err := os.Stat(util.YamlFile); err != nil {
		return errors.New("you need to be in the book directory to run this command (cd <book-name>)")
	}

	return nil
}
