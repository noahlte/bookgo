package util

import (
	"strings"
)

var RestrictedChar = []string{" ", "?", "!", ".", "/", ":", ",", "€", "$", "+", "-", "=", "*", "µ", "¨", "^", "°", "'", "\\", "\"", "<", ">", "|", "#", "%", "&", ";", "`", "@"}

func SanitizeName(name string) string {
	filepath := strings.ToLower(name)
	for _, char := range RestrictedChar {
		filepath = strings.ReplaceAll(filepath, char, "-")
	}

	return filepath
}