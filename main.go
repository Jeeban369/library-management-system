package main

import (
	"fmt"
	"strings"
)

type Book struct {
	Title      string
	Author     string
	Year       int
	IsBorrowed bool
}

type Library struct {
	Books []*Book
}

// adding books
func (lib *Library) addBook(title, author string, year int) {
	book := &Book{
		Title:      title,
		Author:     author,
		Year:       year,
		IsBorrowed: false,
	}
	lib.Books = append(lib.Books, book)
	fmt.Printf("Added book: 📙 \"%s\" by 🖊️ %s (⏳ Year: %d)\n", title, author, year)
}

// remove book by title
func (lib *Library) removeBook(title string) {
	for i, book := range lib.Books {
		if strings.EqualFold(book.Title, title) {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Printf("Remove book: \"%s\"\n", title)
			return
		}
	}
	fmt.Println("Book not found!")
}

// search book by title or author
func (lib *Library) searchBook(query string) {
	fmt.Printf("Searching for: \"%s\"...\n", query)
	found:=false

	for _, book := range lib.Books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)) {
			fmt.Printf("Found: \"%s\" by %s (Year: %d) - Borrowed: %t\n", book.Title, book.Author, book.Year, book.IsBorrowed)
			found= true
		}
	}
	if !found{
		fmt.Println("No books found!")
	}
}

//borrowed or not?
func (lib *Library) borrowBook(title string){
	for _, book := range lib.Books{
		if strings.EqualFold(book.Title, title){
			if !book.IsBorrowed{
				book.IsBorrowed = true
				fmt.Printf("You borrowed \"%s\"\n", book.Title)
			}else{
				fmt.Printf("Book \"%s\" is already borrowed!\n", book.Title)
			}
			return
		}
	}
	fmt.Println("Book not found!")
}

//list of books
func (lib *Library) listBooks(){
	if len(lib.Books) == 0{
		fmt.Println("No books in the library.")
		return
	}
	fmt.Println("library Books:")
	for _, book := range lib.Books{
		fmt.Printf("\"%s\" by %s (Year: %d) - Borrowed: %t\n", book.Title, book.Author, book.Year, book.IsBorrowed)
	}
}

func main() {
	lib := &Library{}

	//added some books
	lib.addBook("The Go Programming Language", "Alan Donovan", 2015)
	lib.addBook("Clean Code", "Robert C. Martin", 2008)
	lib.addBook("Introduction To Algorithms", "Thomas H. Cormen", 2009)

	//list all books
	lib.listBooks()

	//search for a book
	lib.searchBook("Go")

	//borrow a book
	lib.borrowBook("clean code")

	// list all books again to see changes
	lib.listBooks()

	//try to borrow the same book again
	lib.borrowBook("clean code")

	//remove a book
	lib.removeBook("Introduction to Algorithms")

	//list remaining book
	lib.listBooks()

}
