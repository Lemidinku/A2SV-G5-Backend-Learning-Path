package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task3/models"
	"task3/services"
)


type LibraryController struct {
	LibraryService services.LibraryManager
}

func NewLibraryController(l services.LibraryManager) *LibraryController {
	return &LibraryController{LibraryService: l}
}

func (lc *LibraryController) AddBook() {
	reader := bufio.NewReader(os.Stdin)
	var book models.Book
	fmt.Print("Enter Book ID: ")
	var id string
	fmt.Scanln(&id);
	book.ID, _ = strconv.Atoi(id)
	fmt.Print("Enter Book Title: ")
	title, _ := reader.ReadString('\n')
	book.Title = strings.TrimSpace(title)
	fmt.Print("Enter Book Author: ")
	author, _ := reader.ReadString('\n')
	book.Author = strings.TrimSpace(author)
	book.Status = "Available"
	lc.LibraryService.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) RemoveBook() {
	var bookID int
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)
	lc.LibraryService.RemoveBook(bookID)
	fmt.Println("Book removed successfully!")
	
}


func (lc *LibraryController) BorrowBook() {
	var bookID, memberID int
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)
	err := lc.LibraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}


func (lc *LibraryController) ReturnBook() {
	var bookID, memberID int
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&memberID)
	err := lc.LibraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}


func (lc *LibraryController) ListAvailableBooks() {
	books := lc.LibraryService.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("Available Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func (lc *LibraryController) ListBorrowedBooks() {
	var memberID int
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&memberID)
	books := lc.LibraryService.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}


// implement add member
func (lc *LibraryController) AddMember() {
	var member models.Member
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&member.ID)
	fmt.Print("Enter Member Name: ")
	fmt.Scanln(&member.Name)
	lc.LibraryService.AddMember(member)
	fmt.Println("\nMember added successfully!")
}