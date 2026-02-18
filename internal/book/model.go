package book

import (
	"time"
)

type Book struct {
	Name 				string 		`yaml:"name"`
	Description string 		`yaml:"description"`
	Author 			string		`yaml:"author"`
	CreatedAt 	time.Time `yaml:"created-at"`
	Chapters 		[]Chapter `yaml:"chapters"`
}

type Chapter struct {
	Name 				string
	Number 			int
	Description string
	Path 				string
	Sections		[]Section
}

type Section struct {
	Name 		string
	Path 		string
	Content string
}