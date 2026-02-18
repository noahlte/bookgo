package command

import (
	"errors"
	"strings"

	"github.com/noahlte/bookgo/internal/book"
	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

var chapterDescription string

var addChapterCommand = &cobra.Command{
	Use: "add-chapter <name>",
	Short: "Add a new chapter to your book",
	Long: "...",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if args[0] == "" {
			return errors.New("your chapter need a name")
		}

		var name string
		if len(args) > 0 {
			name = strings.Join(args, " ")
		}

		newChapter := &book.Chapter{
			Name: name,
			Description: bookDescription,
		}

		return service.AddChapter(newChapter)
	},
}

func init() {
	rootCmd.AddCommand(addChapterCommand)

	addChapterCommand.Flags().StringVarP(&chapterDescription, "description", "d", "...", "Description of the chapter")
}