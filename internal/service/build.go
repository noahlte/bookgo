package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/util"
	"github.com/playwright-community/playwright-go"
	"github.com/yuin/goldmark"
)



func BuildBook() error {
	var bookpath, err = os.Getwd()
	if err != nil { 
		return err 
	}

	chapters, err := scanContent(bookpath)
	if err != nil { 
		return err 
	}

	content := assembleContent(chapters)
	
	htmlpath, err := convertToHTML(content)
	if err != nil {
		return err
	}
	fmt.Println(htmlpath)

	var userBook book.Book
	err = userBook.UnmarshalBook()
	if err != nil { 
		return err 
	}

	err = convertHTMLtoPDF(htmlpath, bookpath, userBook.Name)
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

/*
scanContent va venir scanner l'entièreté du fichier Content afin de mettre à jour le book.yaml pour qu'il n'y est aucune erreur.
*/
func scanContent(bookpath string) ([]book.Chapter, error) {
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

			sections, err := readSection(chapter.Name(), bookpath)
			if err != nil {
				return nil, err
			}

			newChapter := book.Chapter{
				Name: chapterName,
				Number: index + 1,
				Path: filepath.Join(bookpath, util.ContentDir, chapter.Name()),
				Sections: sections,
			}

			chapters = append(chapters, newChapter)
		}
	}

	return chapters, nil
}

// readSection prend un chapitre du livre et vient scanner tout les fichiers .md afin d'en retourner l'ensemble des sections
func readSection(chapterName string, bookpath string) ([]book.Section, error) {
	sections := make([]book.Section, 0)

	files, err := os.ReadDir(filepath.Join(util.ContentDir, chapterName))
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

		data, err := os.ReadFile(filepath.Join(util.ContentDir, chapterName, section.Name()))
		if err != nil {
			return nil, err
		}

		newSection := book.Section{
			Name: sectionName,
			Path: filepath.Join(bookpath, util.ContentDir, chapterName, section.Name()),
			Content: data,
		}

		sections = append(sections, newSection)
	}

	return sections, nil
}

func assembleContent(chapters []book.Chapter) []byte {
	var finalContent []byte

	for _, chapter := range chapters {
		finalContent = append(finalContent, []byte("# " + chapter.Name + "\n\n")...)
		for _, section := range chapter.Sections {
			finalContent = append(finalContent, section.Content...)
			finalContent = append(finalContent, []byte("\n\n")...)
		}
	}

	return finalContent
}

func convertToHTML(content []byte) (string, error) {
	htmlpath := filepath.Join(os.TempDir(), "book.html")

	f, err := os.OpenFile(htmlpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if err := goldmark.Convert(content, f); err != nil {
		return "", err
	}

	return htmlpath, nil
}

func convertHTMLtoPDF(htmlpath, bookpath, bookname string) error {
	htmlpath = filepath.FromSlash(htmlpath)
	bookname = util.SanitizeName(bookname)

	err := os.MkdirAll("dist", 0755)
	if err != nil {
		return err
	}

	pdfpath := filepath.Join(bookpath, "dist", fmt.Sprintf("%s.pdf", bookname))

	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return err
	}
	defer browser.Close()

	context, err := browser.NewContext()
	if err != nil {
		return err
	}

	page, err := context.NewPage()
	if err != nil {
		return err
	}

	_, err = page.Goto("file:///" + filepath.ToSlash(htmlpath))
	if err != nil {
		return err
	}

	_, err = page.PDF(playwright.PagePdfOptions{
		Path: playwright.String(pdfpath),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Your book has been compiled in %s", pdfpath)

	return nil
}
