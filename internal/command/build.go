package command

import (
	"github.com/noahlte/bookgo/internal/service"
	"github.com/spf13/cobra"
)

var fileSize string

var buildCommand = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Transform your book file into a real PDF book! ",
	Long:    "...",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.BuildBook(fileSize)
	},
}

func init() {
	rootCmd.AddCommand(buildCommand)

	buildCommand.Flags().StringVarP(&fileSize, "file-format", "f", "A4", "Choose the file format of your PDF (A4, A5...)")
}
