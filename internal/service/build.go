package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
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
			bookpath, err := os.Getwd()
			if err != nil {
				return err
			}

			name, ok := strings.CutPrefix(chapter.Name(), prefix)
			if !ok {
				return errors.New("there has been an error while parsing file name")
			}

			name = strings.ReplaceAll(name, "-", " ")
			capitalizeName := strings.Fields(name)
			name = util.Capitalize(capitalizeName)

			/*
			TODO: Nested loop for Section
				- Read chapter dir
				- See if the dir is empty
				- Check each file extension to see .md
				- Analyze content --> to text
				- Convert each file into Section struct
				- Add Section truc to a Section array
				- Add Section array to Chapter
			*/

			newChapter := &book.Chapter{
				Name: name,
				Number: index + 1,
				Path: path.Join(bookpath, util.ContentDir, chapter.Name()),
			}

			fmt.Println(newChapter)
		}
	}

	return nil
}