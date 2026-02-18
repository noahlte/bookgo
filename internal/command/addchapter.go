package command

import (
	"errors"
	"fmt"
	"strings"

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

		var filepath string
		if len(args) > 0 {
			filepath = strings.Join(args, "-")
		}

		fmt.Printf("New chapter %s created at %s", args[0], filepath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addChapterCommand)
}