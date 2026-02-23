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

func CapitalizeWords(name []string) string {
	for i, word := range name {
		name[i] = cases.Title(language.English).String(word)
	}

	return strings.Join(name, " ")
}