package book

import (
	"time"
)

type Book struct {
	name string
	description string
	author string
	createdAt time.Time
}

type Chapter struct {
	name string
	number int
	description string
}