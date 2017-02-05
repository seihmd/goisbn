package goisbn

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

func BenchmarkIsISBN10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsISBN10("157965620X")
		IsISBN10("960425059X")
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

func BenchmarkIsISBN13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsISBN13("9789992158104")
		IsISBN13("9789992158103")
	}
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

func TestConv13To10(t *testing.T) {
	tests := map[string]string{
		"9789992158104": "9992158107",
		"9789971502102": "9971502100",
		"9789604250592": "9604250590",
		"9788090273412": "8090273416",
		"9788535902778": "8535902775",
		"9781843560289": "1843560283",
		"9790205006129": "0205006124",
		"9790205007744": "0205007740",
		"9795000000228": "5000000226",
	}
	for test, expect := range tests {
		isbn10, _ := Conv13To10(test)
		if isbn10 != expect {
			t.Fatal("failed convert isbn10: expect" + expect + " but " + isbn10)
		}
	}
}

func TestConv10To13(t *testing.T) {
	tests := map[string]string{
		"9992158107": "9789992158104",
		"9971502100": "9789971502102",
		"9604250590": "9789604250592",
		"8090273416": "9788090273412",
		"8535902775": "9788535902778",
		"1843560283": "9781843560289",
	}
	for test, expect := range tests {
		isbn13, _ := Conv10To13(test)
		if isbn13 != expect {
			t.Fatal("failed convert isbn13: expect " + expect + " but " + isbn13)
		}
	}
}

func TestExtract(t *testing.T) {
	tests := map[string]string{
		"http://www.shoeisha.co.jp/book/detail/9784798126708": "9784798126708",
		"http://www.oreilly.co.jp/books/9784873117362/":       "9784873117362",
		"http://shop.oreilly.com/product/0636920047124.do":    "0636920047124",
		"http://gihyo.jp/book/2016/9784774185798":             "9784774185798",
		"http://gihyo.jp/book/2006/4774129453":                "4774129453",
		"http://book.impress.co.jp/books/1115101122":          "",
	}

	for test, expect := range tests {
		s, _ := Extract(test)
		if s != expect {
			t.Fatal("failed extract isbn: " + expect + " but " + s)
		}
	}
}

func BenchmarkExtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Extract("http://example.com/9784798126708")
		Extract("http://example.com/9784798126708/abc")
		Extract("http://example.com/097522980X")
		Extract("http://example.com/097522980X/abc")
	}
}

func readFile(filename string) []string {
	file, _ := ioutil.ReadFile(filename)
	contents := string(file)
	ISBNs := strings.Split(string(contents), "\n")
	return ISBNs
}
