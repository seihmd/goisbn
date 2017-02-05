# goisbn
--
    import "github.com/seihmd/goisbn"

Package goisbn contains utility funcs for handling isbn code.

## Usage

#### func  Conv10To13

```go
func Conv10To13(isbn string) (string, error)
```
Conv10To13 convert ISBN10 code to ISBN13.

#### func  Conv13To10

```go
func Conv13To10(isbn string) (string, error)
```
Conv13To10 convert ISBN13 code to ISBN10.
NOTE: having prefix 979 does not have
ISBN10, but can be converted into numerically correct one.

#### func  Extract

```go
func Extract(s string) (string, error)
```
Extract extracts valid isbn code from string (ex. url)

#### func  FormatISBN

```go
func FormatISBN(code string) (string, error)
```
FormatISBN hyphenates code as ISBN10 or 13.
ISBN10: 1234567890 -> 1-23-456789-0
ISBN13: 9781234567890 -> 978-1-23-456789-0

#### func  IsISBN

```go
func IsISBN(code string) bool
```
IsISBN validates code as ISBN. accepts hyphenated or not

#### func  IsISBN10

```go
func IsISBN10(code string) bool
```
IsISBN10 validates code as ISBN10. accepts hyphenated or not

#### func  IsISBN13

```go
func IsISBN13(code string) bool
```
IsISBN13 validates code as ISBN13. accepts hyphenated or not
