package string

import "strings"

func FormatNameString(str *string) *string {
	if str == nil {
		return str
	}

	formattedStr := strings.Title(strings.TrimSpace(*str))

 	return &formattedStr
}

func FormatWhiteSpace(str *string) *string {
	if str == nil {
		return str
	}

	formattedStr := strings.ReplaceAll(strings.TrimSpace(*str), " ", "-")

	return &formattedStr
}

func RemoveEmptyQuote(str *string) *string {
	if str == nil {
		return str
	}

	if *str == "''" {
		formattedStr := ""
		return &formattedStr
	}

	return str
}