package book

import (
	"os"
	"path"
	"time"

	"github.com/noahlte/bookgo/internal/util"
	"gopkg.in/yaml.v3"
)

type Book struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Author      string    `yaml:"author"`
	CreatedAt   time.Time `yaml:"created-at"`
	Chapters    []Chapter `yaml:"chapters"`
}

type Chapter struct {
	Name        string    `yaml:"name"`
	Number      int       `yaml:"number"`
	Description string    `yaml:"description"`
	Path        string    `yaml:"path"`
	Sections    []Section `yaml:"sections"`
}

type Section struct {
	Name    string
	Path    string
	Content string
}

type SectionTemplate struct {
	ChapterName        string
	ChapterDescription string
}

func (b *Book) UnmarshalBook() error {
	f, err := os.ReadFile("book.yaml")
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, &b); err != nil {
		return err
	}

	return nil
}

func (b *Book) Save() error {
	data, err := yaml.Marshal(b)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(util.YamlFile), data, 0644)
	if err != nil {
		return err
	}

	return nil
}
