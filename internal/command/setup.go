package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setupCommand = &cobra.Command{
	Use: "new <name>",
	Short: "Create a new BookGo project",
	Long: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new file...")
	},
}
func init() {
	rootCmd.AddCommand(setupCommand)
}

