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

func BuildBook(fileformat string) error {
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

	err = convertHTMLtoPDF(htmlpath, bookpath, userBook.Name, fileformat)
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

var htmlTemplate = `<!DOCTYPE html>
<html lang="fr">
<head>
<meta charset="UTF-8">
<style>
  body {
    font-family: "Segoe UI", system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
    font-size: 11.5pt;
    line-height: 1.75;
    color: #1a1a1a;
    max-width: 100%;
    margin: 0;
    padding: 0;
  }
  h1 { font-size: 2em; font-weight: 700; margin-top: 2em; margin-bottom: 0.5em; page-break-before: always; }
  h1:first-of-type { page-break-before: avoid; }
  h2 { font-size: 1.35em; font-weight: 600; margin-top: 1.6em; margin-bottom: 0.4em; }
  h3 { font-size: 1.1em; font-weight: 600; margin-top: 1.3em; }
  p  { margin: 0.6em 0; text-align: justify; }
  code {
    font-family: "Cascadia Code", "Fira Code", "Courier New", monospace;
    font-size: 0.85em;
    background: #1e1e2e;
    color: #cdd6f4;
    padding: 0.15em 0.4em;
    border-radius: 4px;
  }
  pre {
    background: #1e1e2e;
    color: #cdd6f4;
    padding: 1.1em 1.2em;
    border-radius: 6px;
    font-size: 0.82em;
    line-height: 1.5;
    overflow-x: auto;
    page-break-inside: avoid;
  }
  pre code { background: none; color: inherit; padding: 0; }
  blockquote {
    border-left: 3px solid #aaa;
    margin: 1em 0;
    padding-left: 1em;
    color: #555;
    font-style: italic;
  }
  hr { border: none; border-top: 1px solid #ddd; margin: 2em 0; }
</style>
</head>
<body>
%s
</body>
</html>`

func convertToHTML(content []byte) (string, error) {
	htmlpath := filepath.Join(os.TempDir(), "book.html")

	var buf strings.Builder
	if err := goldmark.Convert(content, &buf); err != nil {
		return "", err
	}

	full := strings.Replace(htmlTemplate, "%s", buf.String(), 1)

	if err := os.WriteFile(htmlpath, []byte(full), 0644); err != nil {
		return "", err
	}

	return htmlpath, nil
}

func convertHTMLtoPDF(htmlpath, bookpath, bookname, fileformat string) error {
	displayName := bookname
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

	htmlURL := "file:///" + filepath.ToSlash(htmlpath)
	fmt.Println("URL:", htmlURL)
	_, err = page.Goto(htmlURL)
	if err != nil {
		return err
	}

	headerHTML := fmt.Sprintf(
		`<div style="font-size:8px;width:100%%;text-align:center;color:#aaa;padding-top:6px;">%s</div>`,
		displayName,
	)
	footerHTML := `<div style="font-size:8px;width:100%;text-align:center;color:#aaa;padding-bottom:6px;">` +
		`<span class="pageNumber"></span> / <span class="totalPages"></span></div>`

	_, err = page.PDF(playwright.PagePdfOptions{
		Path:                playwright.String(pdfpath),
		Format:              playwright.String(fileformat),
		PrintBackground:     playwright.Bool(true),
		DisplayHeaderFooter: playwright.Bool(true),
		HeaderTemplate:      playwright.String(headerHTML),
		FooterTemplate:      playwright.String(footerHTML),
		Margin: &playwright.Margin{
			Top:    playwright.String("18mm"),
			Bottom: playwright.String("18mm"),
			Left:   playwright.String("15mm"),
			Right:  playwright.String("15mm"),
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Your book has been compiled in %s", pdfpath)

	return nil
}
