package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrStringContainsNumbers = errors.New("string contains numbers")
	ErrInvalidString         = errors.New("invalid string")
)

func Unpack(stringForUnpacking string) (string, error) {
	if stringForUnpacking == "" {
		return "", nil
	}
	runes := []rune(stringForUnpacking)
	if !unicode.IsLetter(runes[0]) {
		return "", ErrInvalidString
	}

	var result strings.Builder
	for i := 0; i < len(runes); i++ {
		current := runes[i]
		next := rune(0)
		if i+1 < len(runes) {
			next = runes[i+1]
		}
		if unicode.IsLetter(current) {
			if unicode.IsDigit(next) {
				repeatCount := int(next - '0')
				result.WriteString(strings.Repeat(string(current), repeatCount))
			} else {
				result.WriteString(string(current))
			}
		} else if unicode.IsDigit(current) && unicode.IsDigit(next) {
			return result.String(), ErrStringContainsNumbers
		}
	}

	return result.String(), nil
}
