package service

import "fmt"

// TODO: Setup command service
func SetupBook(name, author string) error {
	fmt.Printf("Creating a new book : %s writed by %s", name, author)

	return nil
}