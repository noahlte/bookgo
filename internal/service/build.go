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
	chapters, err := scanContent()
	if err != nil { 
		return err 
	}

	var userBook book.Book
	err = userBook.UnmarshalBook()
	if err != nil { 
		return err 
	}

	userBook.Chapters = chapters

	err = userBook.Save()
	if err != nil { 
		return err 
	}

	return nil
}


func scanContent() ([]book.Chapter, error) {
	bookpath, err := os.Getwd()
	if err != nil { 
		return nil, err 
	}

	if _, err := os.Stat("content"); errors.Is(err, fs.ErrNotExist) {
		return nil, errors.New("no content directory found")
	}

	folders, err := os.ReadDir("content")
	if err != nil {
		return nil, errors.New("failed to read directory")
	}

	if len(folders) <= 0 {
		return nil, errors.New("content directory is empty")
	}

	chapters := make([]book.Chapter, 0)

	for index, chapter := range folders {
		if chapter.IsDir() {
			prefix := fmt.Sprintf("%d-chapter-", index + 1)

			chapterName, ok := strings.CutPrefix(chapter.Name(), prefix)
			if !ok { 
				return nil, errors.New("there has been an error while parsing file name") 
			}

			chapterName = strings.ReplaceAll(chapterName, "-", " ")
			chapterWords := strings.Fields(chapterName)
			chapterName = util.Capitalize(chapterWords)

			sections := make([]book.Section, 0)

			files, err := os.ReadDir(path.Join(util.ContentDir, chapter.Name()))
			if err != nil { 
				return nil, err 
			}
			if len(files) == 0 { 
				return nil, errors.New("a chapter can't be empty") 
			}

			for _, section := range files {
				if !strings.HasSuffix(section.Name(), ".md") { continue }

				sectionName, ok := strings.CutSuffix(section.Name(), ".md")
				if !ok { 
					return nil, errors.New("can not cut suffix") 
				}

				sectionName = strings.ReplaceAll(sectionName, "-", " ")
				sectionWords := strings.Fields(sectionName)
				sectionName = util.Capitalize(sectionWords)

				// TODO: Scan content

				newSection := &book.Section{
					Name: sectionName,
					Path: path.Join(bookpath, util.ContentDir, chapter.Name(), section.Name()),
				}

				sections = append(sections, *newSection)
			}


			newChapter := &book.Chapter{
				Name: chapterName,
				Number: index + 1,
				Path: path.Join(bookpath, util.ContentDir, chapter.Name()),
				Sections: sections,
			}

			chapters = append(chapters, *newChapter)
		}
	}

	return chapters, nil
}