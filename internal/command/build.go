package command

import (
	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

var buildCommand = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Transform your book file into a real PDF book! ",
	Long:    "...",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.BuildBook()
	},
}

func init() {
	rootCmd.AddCommand(buildCommand)
}
