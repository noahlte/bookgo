package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/util"
)

func BuildBook() error {
	err := scanContent()
	if err != nil {
		return err
	}

	return nil
}


func scanContent() error {
	if _, err := os.Stat("content"); errors.Is(err, fs.ErrNotExist) {
		return errors.New("no content directory found")
	}

	chapters, err := os.ReadDir("content")
	if err != nil {
		return errors.New("failed to read directory")
	}

	if len(chapters) <= 0 {
		return errors.New("content directory is empty")
	}

	for index, chapter := range chapters {
		if chapter.IsDir() {
			prefix := fmt.Sprintf("%d-chapter-", index + 1)

			name, ok := strings.CutPrefix(chapter.Name(), prefix)
			if !ok {
				return errors.New("there has been an error while parsing file name")
			}

			name = strings.ReplaceAll(name, "-", " ")
			capitalizeName := strings.Fields(name)
			name = util.CapitalizeWords(capitalizeName)

			newChapter := &book.Chapter{
				Name: name,
				Number: index + 1,
			}

			fmt.Println(newChapter)
		}
	}

	return nil
}