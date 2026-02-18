package command

import (
	"errors"

	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

var author string

var setupCommand = &cobra.Command{
	Use: "new <name> <author>",
	Short: "Create a new BookGo project",
	Long: "...",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if name == "" {
			return errors.New("You can't create a project without a name (no space)")
		}

		return service.SetupBook(name, author)
	},
}

func init() {
	rootCmd.AddCommand(setupCommand)

	setupCommand.Flags().StringVarP(&author, "author", "a", "John Doe", "Author name")
}

