package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
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

	for _, chapter := range chapters {
		fmt.Println(chapter)
	}

	return nil
}