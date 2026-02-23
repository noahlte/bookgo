package util

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var RestrictedChar = []string{" ", "?", "!", ".", "/", ":", ",", "€", "$", "+", "-", "=", "*", "µ", "¨", "^", "°", "'", "\\", "\"", "<", ">", "|", "#", "%", "&", ";", "`", "@"}

func SanitizeName(name string) string {
	filepath := strings.ToLower(name)
	for _, char := range RestrictedChar {
		filepath = strings.ReplaceAll(filepath, char, "-")
	}

	return filepath
}

func Capitalize(name []string) string {
	name[0] = cases.Title(language.English).String(name[0])

	return strings.Join(name, " ")
}