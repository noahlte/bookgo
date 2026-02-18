package util

import (
	"strings"
)

func SanitizeName(name string) string {
	filepath := strings.ToLower(name)
	for _, char := range RestrictedChar {
		filepath = strings.ReplaceAll(filepath, char, "-")
	}

	return filepath
}