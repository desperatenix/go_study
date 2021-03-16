package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var prev rune
	var b strings.Builder
	errPattern := regexp.MustCompile(`(^\d|\d{2,})`)
	if errPattern.FindString(s) != "" {
		return "", ErrInvalidString
	}
	if s == "" {
		return "", nil
	}
	for _, cur := range s {
		if unicode.IsDigit(cur) && string(cur) != "0" { //nolint:gocritic
			integer, _ := strconv.Atoi(string(cur))
			b.WriteString(strings.Repeat(string(prev), integer-1))
		} else if string(cur) == "0" {
			crop := b.String()[:b.Len()-1]
			b.Reset()
			b.WriteString(crop)
		} else {
			b.WriteString(string(cur))
		}
		prev = cur
	}
	return b.String(), nil
}
