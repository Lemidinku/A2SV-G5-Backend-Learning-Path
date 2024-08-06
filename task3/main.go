package main

import (
	"fmt"
	"strconv"
	"task3/controllers"
	"task3/models"
	"task3/services"
)

func main() {
	library := services.Library{Books: make(map[int]models.Book), Members: make(map[int]models.Member)}
	controller := controllers.NewLibraryController(&library)

	actions := map[int]func(){
		1: controller.AddBook,
		2: controller.RemoveBook,
		3: controller.BorrowBook,
		4: controller.ReturnBook,
		5: controller.ListAvailableBooks,
		6: controller.ListBorrowedBooks,
		7: controller.AddMember,
	}
	fmt.Println("___________Welcome to the Library Management System!_________")
	// if the library is empty, add some books and a member

	if len(library.Books) == 0 {
		fmt.Println("The library is empty. Let's add some books and a member.")
		controller.AddBook()
		controller.AddMember()
	}

	for {
		fmt.Println("\nChoose an action: (ex. type \"1\" to add a book)")
		choices := []string{
		"Add a book",
		"Remove a book",
		"Borrow a book",
		"Return a book",
		"List Available books",
		"List borrowed books",
		"Add a member",
	}
		for idx, choice:= range choices {
			fmt.Printf("\t %d. %s \n", idx+1, choice)
		}
		fmt.Println()
		var action string
		fmt.Scanln(&action)
		action_num, _ := strconv.Atoi(action)
		if f, ok := actions[action_num]; ok {
			fmt.Println()
			f()
		} else if action != "" {
			fmt.Println("Invalid action!")
		}
	}

	}