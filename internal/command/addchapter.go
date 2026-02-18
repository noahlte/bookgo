package command

import (
	"errors"
	"strings"

	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

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

		return service.AddChapter(name)
	},
}

func init() {
	rootCmd.AddCommand(addChapterCommand)
}