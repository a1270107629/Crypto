package tools

import "strings"

func trimString(s string) string {

	return strings.ReplaceAll(s, " ", "")
}
