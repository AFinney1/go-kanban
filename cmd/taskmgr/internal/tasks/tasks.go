//  Package for managing tasks (adding, listing, updating, etc).

package tasks

import (
	"fmt"
	"go-kanban/cmd/taskmgr/internal/utils"
)

// Board represents a Kanban board with 3 columns of cards
type Board struct {
	Todo       []Card // Slice of cards in the Todo column
	InProgress []Card // Slice of cards in the In Progress column
	Done       []Card // Slice of cards in the Done column
}

// Card represents a single card on the board
type Card struct {
	Id          int    // Unique ID of the card
	Title       string // Title of the card
	Description string // Description of the task
	Priority    int    // Priority of the task (higher is more important)
}

// CreateBoard creates a new empty Kanban board
func CreateBoard() Board {
	return Board{
		Todo:       []Card{}, // Empty Todo column
		InProgress: []Card{}, // Empty In Progress column
		Done:       []Card{}, // Empty Done column
	}
}

func CreateCard() Card {
	return Card{
		Id:          utils.RandomInt(1, 100000000),
		Title:       "",
		Description: "",
		Priority:    0,
	}
}

func (b *Board) Print() {
	fmt.Println("[Todo]")
	for _, card := range b.Todo {
		fmt.Printf("  • %s %d\n", card.Title, card.Priority)
		fmt.Printf("    -   %s\n", card.Description)
	}

	fmt.Println("\n[InProgress]")
	for _, card := range b.InProgress {
		fmt.Printf("  • %s %d \n", card.Title, card.Priority)
		fmt.Printf("    -   %s\n", card.Description)
	}

	fmt.Println("\n[Done]")
	//fmt.Println("\n")
	for _, card := range b.Done {
		fmt.Printf("  • %s %d\n", card.Title, card.Priority)
		fmt.Printf("    -   %s\n", card.Description)
	}
}

// AddCardToColumn adds a card to the appropriate column on the board
func (b *Board) AddCardToColumn(column string, card Card) {
	// Switch on the column name and append the card to the appropriate column slice
	switch column {
	case "Todo":
		b.Todo = append(b.Todo, card)
	case "InProgress":
		b.InProgress = append(b.InProgress, card)
	case "Done":
		b.Done = append(b.Done, card)
	}
}

// MoveCard moves a card from one column to another
func (b *Board) MoveCard(source, dest string, title string) {
	// Find the card with the given ID in the source column
	var cardToMove Card
	for _, card := range b.Todo {
		if card.Title == title {
			cardToMove = card
		}
	}
	id := cardToMove.Id

	// Remove the card from the source column
	switch source {
	case "Todo":
		b.Todo = remove(b.Todo, id)
	case "In Progress":
		b.InProgress = remove(b.InProgress, id)
	}

	// Add the card to the destination column
	switch dest {
	case "Todo":
		b.Todo = append(b.Todo, cardToMove)
	case "In Progress":
		b.InProgress = append(b.InProgress, cardToMove)
	case "Done":
		b.Done = append(b.Done, cardToMove)
	}
}

// UpdateCard updates a card on the board
func (b *Board) UpdateCard(id int, updates Card) {
	// Loop through each column and update the card with the given ID
	for i := range b.Todo {
		if b.Todo[i].Id == id {
			b.Todo[i] = updates
		}
	}

	for i := range b.InProgress {
		if b.InProgress[i].Id == id {
			b.InProgress[i] = updates
		}
	}

	for i := range b.Done {
		if b.Done[i].Id == id {
			b.Done[i] = updates
		}
	}
}

// RemoveCard removes a card from the board
func (b *Board) RemoveCard(id int) {
	b.Todo = remove(b.Todo, id)             // Remove from Todo
	b.InProgress = remove(b.InProgress, id) // Remove from In Progress
	b.Done = remove(b.Done, id)             // Remove from Done
}

// remove is a helper function to remove a card from a slice by ID
func remove(slice []Card, id int) []Card {
	for i := range slice {
		if slice[i].Id == id {
			return append(slice[:i], slice[i+1:]...) // Return slice without the card at index i
		}
	}
	return slice // If the card wasn't found, return the original slice
}
