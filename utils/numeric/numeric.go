package numeric

import "regexp"

func IsNumeric(word string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(word)
}
