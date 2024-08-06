package services

import (
	"errors"
	"fmt"
	. "task3/models"
)

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
	AddMember(member Member)
}

type Library struct {
	Books   map[int]Book
	Members map[int]Member
}

// Implement the AddBook method for the Library struct.
func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

// Implement the RemoveBook method for the Library struct.
func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}


// BorrowBook borrows a book from the library for a specific member.
// It takes the bookID and memberID as parameters and returns an error if any.
// If the book is not found in the library, it returns an error "book not found".
// If the book is already borrowed, it returns an error "book is already borrowed".
// If the member is not found in the library, it returns an error "member not found".
// Otherwise, it updates the book status to "Borrowed" and appends the book to the member's borrowed books list.
// It then updates the book and member information in the library and returns nil.
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}
	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Books[bookID] = book
	l.Members[memberID] = member
	fmt.Println(member.BorrowedBooks)
	return nil
}


// Implement the ReturnBook method for the Library struct.
// ReturnBook returns a borrowed book to the library.
// It takes the bookID and memberID as parameters and returns an error if any.
// If the book is not found in the library, it returns an error with the message "book not found".
// If the book is available (not borrowed), it returns an error with the message "book is not borrowed".
// If the member is not found in the library, it returns an error with the message "member not found".
// Otherwise, it updates the book's status to "Available" and removes the book from the member's borrowed books list.
// It then updates the library's book and member records accordingly.
// Returns nil if the book is successfully returned.
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}
	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	book.Status = "Available"
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	l.Books[bookID] = book
	l.Members[memberID] = member
	return nil
}


// Implement the ListAvailableBooks method for the Library struct.
func (l *Library) ListAvailableBooks() []Book {
	var availableBooks []Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// Implement the ListBorrowedBooks method for the Library struct.
func (l *Library) ListBorrowedBooks(memberID int) []Book {
	member, ok := l.Members[memberID]
	if !ok {
		return nil
	}
	return member.BorrowedBooks
}

// Implement the AddMember method for the Library struct.
func (l *Library) AddMember(member Member) {
	l.Members[member.ID] = member
}

