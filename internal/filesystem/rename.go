package filesystem

import (
	"strings"

	"github.com/noahlte/bookgo/internal/util"
)

func RenameFile(name string) string {
	filepath := strings.ToLower(name)
	for _, char := range util.RestrictedChar {
		filepath = strings.ReplaceAll(filepath, char, "-")
	}

	return filepath
}