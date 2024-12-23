# Library Management System

A simple **Library Management System** implemented in Go using **structs** and **pointers**. This program allows users to manage a collection of books by adding, removing, searching, and borrowing books from the library.

---

## Features

1. **Add Books**: Add a new book to the library with its title, author, and publication year.
2. **Remove Books**: Remove a book from the library by its title.
3. **Search Books**: Search for books in the library by title or author.
4. **Borrow Books**: Mark a book as borrowed.
5. **List Books**: Display all books in the library with their details, including whether they are borrowed.

---

## Code Overview

### Book Struct
The `Book` struct represents a single book in the library with the following fields:
- `Title` (string): Title of the book
- `Author` (string): Author of the book
- `Year` (int): Year of publication
- `IsBorrowed` (bool): Indicates whether the book is borrowed

### Library Struct
The `Library` struct contains:
- `Books` ([]*Book): A slice of pointers to `Book` structs, allowing direct modification of the book details.

### Methods

#### `addBook`
Adds a new book to the library.
```go
func (lib *Library) addBook(title, author string, year int)
```

#### `removeBook`
Removes a book from the library by title.
```go
func (lib *Library) removeBook(title string)
```

#### `searchBook`
Searches for books by title or author.
```go
func (lib *Library) searchBook(query string)
```

#### `borrowBook`
Marks a book as borrowed.
```go
func (lib *Library) borrowBook(title string)
```

#### `listBooks`
Lists all books currently in the library.
```go
func (lib *Library) listBooks()
```

---

## How to Run

1. **Install Go**: Ensure you have Go installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

2. **Clone the Repository**: Copy the code into a `.go` file, for example, `library.go`.

3. **Run the Program**:
```bash
go run library.go
```

---

## Example Output

```
Added book: 📙 "The Go Programming Language" by 🖊️ Alan Donovan (⏳ Year: 2015)
Added book: 📙 "Clean Code" by 🖊️ Robert C. Martin (⏳ Year: 2008)
Added book: 📙 "Introduction To Algorithms" by 🖊️ Thomas H. Cormen (⏳ Year: 2009)

Library Books:
"The Go Programming Language" by Alan Donovan (Year: 2015) - Borrowed: false
"Clean Code" by Robert C. Martin (Year: 2008) - Borrowed: false
"Introduction To Algorithms" by Thomas H. Cormen (Year: 2009) - Borrowed: false

Searching for: "Go"...
Found: "The Go Programming Language" by Alan Donovan (Year: 2015) - Borrowed: false

You borrowed "Clean Code"

Library Books:
"The Go Programming Language" by Alan Donovan (Year: 2015) - Borrowed: false
"Clean Code" by Robert C. Martin (Year: 2008) - Borrowed: true
"Introduction To Algorithms" by Thomas H. Cormen (Year: 2009) - Borrowed: false

Book "Clean Code" is already borrowed!

Removed book: "Introduction To Algorithms"

Library Books:
"The Go Programming Language" by Alan Donovan (Year: 2015) - Borrowed: false
"Clean Code" by Robert C. Martin (Year: 2008) - Borrowed: true
```



