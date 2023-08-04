package tasks

import "testing"

func TestCreateBoard(t *testing.T) {
	board := CreateBoard()

	if board.Todo == nil {
		t.Error("Todo column is nil")
	}

	if board.InProgress == nil {
		t.Error("In Progress column is nil")
	}

	if board.Done == nil {
		t.Error("Done column is nil")
	}
}

func TestAddCardToColumn(t *testing.T) {
	board := CreateBoard()

	card := Card{
		Id:          1,
		Title:       "Test card",
		Description: "This is a test card",
		Priority:    2,
	}

	board.AddCardToColumn("Todo", card)

	if len(board.Todo) != 1 {
		t.Errorf("Expected Todo column to have 1 card, got %d", len(board.Todo))
	}

	if board.Todo[0].Id != 1 {
		t.Errorf("Expected card ID to be 1, got %d", board.Todo[0].Id)
	}
}

func TestMoveCard(t *testing.T) {
	board := CreateBoard()

	card := Card{
		Id:          1,
		Title:       "Test card",
		Description: "This is a test card",
		Priority:    2,
	}

	board.AddCardToColumn("Todo", card)
	board.MoveCard("Todo", "In Progress", "test_card")

	if len(board.Todo) != 0 {
		t.Errorf("Expected Todo column to be empty, got %d cards", len(board.Todo))
	}

	if len(board.InProgress) != 1 {
		t.Errorf("Expected In Progress column to have 1 card, got %d", len(board.InProgress))
	}

	if board.InProgress[0].Id != 1 {
		t.Errorf("Expected card ID to be 1, got %d", board.InProgress[0].Id)
	}
}

// Additional tests...
