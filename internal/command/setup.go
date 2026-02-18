package command

import (
	"errors"
	"strings"

	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

var author string

var setupCommand = &cobra.Command{
	Use: "new <name>",
	Short: "Create a new BookGo project",
	Long: "...",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if name == "" {
			return errors.New("You can't create a project without a name (no space)")
		}

		var filepath string
		if len(args) > 0 {
			filepath = strings.Join(args, "-")
		}

		return service.SetupBook(name, author, filepath)
	},
}

func init() {
	rootCmd.AddCommand(setupCommand)

	setupCommand.Flags().StringVarP(&author, "author", "a", "John Doe", "Author name")
}

