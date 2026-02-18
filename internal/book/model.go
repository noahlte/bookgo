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
	Name 				string 		`yaml:"name"`
	Number 			int				`yaml:"number"`
	Description string		`yaml:"description"`
	Path 				string		`yaml:"path"`
	Sections		[]Section	`yaml:"sections"`
}

type Section struct {
	Name 		string
	Path 		string
	Content string
}