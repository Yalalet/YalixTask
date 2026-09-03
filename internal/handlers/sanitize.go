package handlers

import "html"

func inputSanitization(input string) string {
	return html.EscapeString(input)
}
