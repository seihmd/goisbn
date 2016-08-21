package isbn

import (
	"io/ioutil"
	"strings"
	"testing"
)

const (
	testFileDir           = "testfiles"
	testfileISBN10Valid   = testFileDir + "/isbn10_valid.txt"
	testfileISBN10Invalid = testFileDir + "/isbn10_invalid.txt"
	testfileISBN13Valid   = testFileDir + "/isbn13_valid.txt"
	testfileISBN13Invalid = testFileDir + "/isbn13_invalid.txt"
)

func TestIsISBN(t *testing.T) {
	validISBN10s := readFile(testfileISBN10Valid)
	for _, isbn := range validISBN10s {
		result := IsISBN(isbn)
		if !result {
			t.Fatal("Should be valid ISBN10: " + isbn)
		}
	}

	invalidISBN10s := readFile(testfileISBN10Invalid)
	for _, isbn := range invalidISBN10s {
		result := IsISBN(isbn)
		if result {
			t.Fatal("Should be invalid ISBN10: " + isbn)
		}
	}

	validISBN13s := readFile(testfileISBN13Valid)
	for _, isbn := range validISBN13s {
		result := IsISBN(isbn)
		if !result {
			t.Fatal("Should be valid ISBN13: " + isbn)
		}
	}

	invalidISBN13s := readFile(testfileISBN13Invalid)
	for _, isbn := range invalidISBN13s {
		result := IsISBN(isbn)
		if result {
			t.Fatal("Should be invalid ISBN13: " + isbn)
		}
	}
}

func TestIsISBN10(t *testing.T) {
	validISBNs := readFile(testfileISBN10Valid)
	for _, isbn := range validISBNs {
		result := IsISBN10(isbn)
		if !result {
			t.Fatal("Should be valid ISBN10: " + isbn)
		}
	}

	invalidISBNs := readFile(testfileISBN10Invalid)
	for _, isbn := range invalidISBNs {
		result := IsISBN10(isbn)
		if result {
			t.Fatal("Should be invalid ISBN10: " + isbn)
		}
	}
}

func TestIsISBN13(t *testing.T) {
	validISBNs := readFile(testfileISBN13Valid)
	for _, isbn := range validISBNs {
		result := IsISBN13(isbn)
		if !result {
			t.Fatal("Should be valid ISBN13: " + isbn)
		}
	}

	invalidISBNs := readFile(testfileISBN13Invalid)
	for _, isbn := range invalidISBNs {
		result := IsISBN13(isbn)
		if result {
			t.Fatal("Should be invalid ISBN13: " + isbn)
		}
	}
}

func readFile(filename string) []string {
	file, _ := ioutil.ReadFile(filename)
	contents := string(file)
	ISBNs := strings.Split(string(contents), "\n")
	return ISBNs
}

func TestFormatISBN10(t *testing.T) {
	isbn10 := "1234567890"
	formatted := formatISBN10(isbn10)
	if formatted != "1-23-456789-0" {
		t.Fatal("wrong format as isbn10: " + formatted)
	}
}

func TestFormatISBN13(t *testing.T) {
	isbn13 := "9789992158104"
	formatted := formatISBN13(isbn13)
	if formatted != "978-9-99-215810-4" {
		t.Fatal("wrong format as isbn13: " + formatted)
	}
}
