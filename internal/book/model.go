package book

import (
	"time"
)

type Book struct {
	Name 				string
	Description string
	Author 			string
	CreatedAt 	time.Time
	Chapters 		[]Chapter
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