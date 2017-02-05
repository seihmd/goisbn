// Package goisbn contains utility funcs for handling isbn code.
package goisbn

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	isbn10Length  = 10
	isbn13Length  = 13
	isbn10Pattern = "[0-9]{9}[0-9xX]"
	isbn13Pattern = "[0-9]{13}"
)

var regISBN10 = regexp.MustCompile(isbn10Pattern)
var regISBN13 = regexp.MustCompile(isbn13Pattern)

// IsISBN validates code as ISBN. accepts hyphenated or not
func IsISBN(code string) bool {
	switch len(code) {
	case isbn10Length:
		{
			return IsISBN10(code)
		}
	case isbn13Length:
		{
			return IsISBN13(code)
		}
	}
	return false
}

// IsISBN10 validates code as ISBN10. accepts hyphenated or not
func IsISBN10(code string) bool {
	code = removeHyphen(code)
	if len(code) != isbn10Length {
		return false
	}
	var sum int
	for i, r := range code {
		digitStr := string(r)
		var digit int
		if strings.ToLower(digitStr) == "x" {
			digit = 10
		} else {
			digit, _ = strconv.Atoi(digitStr)
		}
		sum += digit * (10 - i)
	}
	return sum%11 == 0
}

// IsISBN13 validates code as ISBN13. accepts hyphenated or not
func IsISBN13(code string) bool {
	code = removeHyphen(code)
	if len(code) != isbn13Length {
		return false
	}
	var sum int
	for i, r := range code {
		digit, _ := strconv.Atoi(string(r))
		multiply := 1
		if i%2 != 0 {
			multiply = 3
		}
		sum += digit * multiply
	}
	return sum%10 == 0
}

// FormatISBN hyphenates code as ISBN10 or 13.
// ISBN10: 1234567890 -> 1-23-456789-0
// ISBN13: 9781234567890 -> 978-1-23-456789-0
func FormatISBN(code string) (string, error) {
	if len(code) == 10 {
		return formatISBN10(code), nil
	} else if len(code) == 13 {
		return formatISBN13(code), nil
	}
	return "", errors.New("invalid length")
}

func formatISBN10(isbn10 string) string {
	return fmt.Sprintf("%s-%s-%s-%s", string(isbn10[0]), string(isbn10[1:3]), string(isbn10[3:9]), string(isbn10[9]))
}

func formatISBN13(isbn13 string) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", string(isbn13[0:3]), string(isbn13[3]), string(isbn13[4:6]), string(isbn13[6:12]), string(isbn13[12]))
}

// Conv10To13 convert ISBN10 code to ISBN13.
func Conv10To13(isbn string) (string, error) {
	if !IsISBN10(isbn) {
		return "", errors.New(isbn + "IS NOT ISBN10")
	}
	// NOTE: isbn13 has 979 prefix has no isbn10 code.
	baseString := "978" + isbn[0:9]
	var sum int
	for i, r := range baseString {
		digit, _ := strconv.Atoi(string(r))
		multiply := 1
		if i%2 != 0 {
			multiply = 3
		}
		sum += digit * multiply
	}
	checkdigit := 10 - sum%10
	if checkdigit == 10 {
		checkdigit = 0
	}
	return baseString + strconv.Itoa(checkdigit), nil
}

// Conv13To10 convert ISBN13 code to ISBN10.
// NOTE: having prefix 979 does not have ISBN10, but can be converted into numerically correct one.
func Conv13To10(isbn string) (string, error) {
	if !IsISBN13(isbn) {
		return "", errors.New(isbn + "IS NOT ISBN13")
	}
	baseString := isbn[3:12]
	var sum int
	for i, r := range baseString {
		digit, _ := strconv.Atoi(string(r))
		sum += digit * (10 - i)
	}
	checkdigit := strconv.Itoa(11 - sum%11)
	if checkdigit == "10" {
		checkdigit = "X"
	}
	if checkdigit == "11" {
		checkdigit = "0"
	}
	return baseString + checkdigit, nil
}

// Extract extracts valid isbn code from string (ex. url)
func Extract(s string) (string, error) {
	match := regISBN13.FindStringSubmatch(s)
	if len(match) == 0 {
		match = regISBN10.FindStringSubmatch(s)
		if len(match) == 0 {
			return "", errors.New("ISBN not found: " + s)
		}
	}
	if !IsISBN(match[0]) {
		return "", errors.New("ISBN not found: " + s)
	}
	return match[0], nil
}

func removeHyphen(s string) string {
	return strings.Replace(s, "-", "", -1)
}
