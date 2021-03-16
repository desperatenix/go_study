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
		switch {
		case unicode.IsDigit(cur) && string(cur) != "0":
			integer, _ := strconv.Atoi(string(cur))
			b.WriteString(strings.Repeat(string(prev), integer-1))
		case string(cur) == "0":
			crop := b.String()[:b.Len()-1]
			b.Reset()
			b.WriteString(crop)
		default:
			b.WriteString(string(cur))
		}
		prev = cur
	}
	return b.String(), nil
}
