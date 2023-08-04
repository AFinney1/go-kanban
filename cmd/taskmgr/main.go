// Contains the main entry point (main.go) which will initialize the CLI app and start the prompt/REPL.

package main

import (
	"fmt"
	"go-kanban/cmd/taskmgr/internal/tasks"
	"go-kanban/cmd/taskmgr/internal/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Parse CLI args
	//args := os.Args[1:]

	// Create a new board
	board := tasks.CreateBoard()

	// Start REPL loop
	for {
		// Prompt user for input
		fmt.Println("\n \n ------------------------------Kanban Gopher------------------------------")
		fmt.Println("Enter a command: \n - add (Ex. add 'New card title' 'Column title' 'Description')\n - move (Ex. move 'Card title' 'Source column' 'Destination column') \n - update (Ex. update 'New title' 'New contents')\n - remove (Ex. remove 'title') \n - quit   \n   \n \n ")

		board.Print()
		input := utils.GetUserInput("> ")

		// Split input into words
		words := strings.Fields(input)
		command := words[0]

		switch command {
		case "add":
			// Add a new card
			// Usage: add "New card title" "Column title"
			column := words[1]
			title := words[2]
			description := strings.Join(words[3:], " ")
			card := tasks.Card{
				Id:          utils.RandomInt(1, 100000000),
				Title:       title,
				Description: description,
				Priority:    1,
			}
			board.AddCardToColumn(column, card) // title could just be "Todo	"
			//board.Print()

		case "move":
			// Move a card to a new column
			// Usage: move "Card title" "New column"
			source := words[1]
			cardTitle := words[2]
			newColumn := words[3]
			// source, dest, cardtitle
			board.MoveCard(source, newColumn, cardTitle)
			//board.Print()

		case "update":
			// Update a card's contents
			// Usage: update ID "New title" "New contents"
			id, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Println("An error occured while parsing the ID")
				fmt.Println(err)
			}
			newTitle := words[2]
			newContents := strings.Join(words[3:], " ")
			newCard := tasks.Card{Id: id, Title: newTitle, Description: newContents, Priority: 1}
			board.UpdateCard(id, newCard)
			//board.Print()

		case "remove":
			// Remove a card from the board
			// Usage: remove ID
			id, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Println("An error occured while parsing the ID")
				fmt.Println(err)
			}
			board.RemoveCard(id)
			//board.Print()

		case "quit":
			// Exit the REPL
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try add, move, update or quit.")
		}
	}
}
