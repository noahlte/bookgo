package command

import (
	"errors"
	"strings"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

var author string
var bookDescription string

var setupCommand = &cobra.Command{
	Use:   "new <name>",
	Short: "Create a new BookGo project",
	Long:  "...",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if name == "" {
			return errors.New("you can't create a project without a name (no space)")
		}

		if len(args) > 0 {
			name = strings.Join(args, " ")
		}

		newBook := &book.Book{
			Name:        name,
			Author:      author,
			Description: bookDescription,
		}

		return service.SetupBook(newBook)
	},
}

func init() {
	rootCmd.AddCommand(setupCommand)

	setupCommand.Flags().StringVarP(&author, "author", "a", "John Doe", "Author name")
	setupCommand.Flags().StringVarP(&bookDescription, "description", "d", "...", "Description of the book")
}
