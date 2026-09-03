package service

import "regexp"

var htmlTagRegex = regexp.MustCompile(`<[^>]*>`)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func containsHTML(textInput string) bool {
	return htmlTagRegex.MatchString(textInput)
}
