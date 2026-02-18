package filesystem

import (
	"errors"
	"os"

	"github.com/noahlte/bookgo/internal/util"
)

func FindBookRoot() error {
	if _, err := os.Stat(util.YamlFile); err != nil {
		return errors.New("you need to be in the book directory to run this command (cd <book-name>)")
	}

	return nil
}