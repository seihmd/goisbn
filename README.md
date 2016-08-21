# isbn
--
    import "github.com/seihmd/go-isbn"

Package isbn contains utility funcs for handling isbn code.

## Usage

#### func  Conv10To13

```go
func Conv10To13(isbn string) (string, error)
```

Conv10To13 convert ISBN10 code to ISBN13

#### func  Conv13To10

```go
func Conv13To10(isbn string) (string, error)
```

Conv13To10 convert ISBN13 code to ISBN10

#### func  FormatISBN

```go
func FormatISBN(code string) (string, error)
```

FormatISBN hyphenates code as ISBN10 or 13

#### func  IsISBN

```go
func IsISBN(code string) bool
```

IsISBN validates code as ISBN

#### func  IsISBN10

```go
func IsISBN10(code string) bool
```

IsISBN10 validates code as ISBN10

#### func  IsISBN13

```go
func IsISBN13(code string) bool
```

IsISBN13 validates code as ISBN13
