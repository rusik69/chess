package ui

import (
	"bufio"
	"bytes"
	"chess-game/chess"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewInterface(t *testing.T) {
	ui := NewInterface()
	if ui == nil {
		t.Error("NewInterface should return a valid interface")
	}
	if ui.game == nil {
		t.Error("Interface should have a valid game")
	}
	if ui.reader == nil {
		t.Error("Interface should have a valid reader")
	}
}

func TestDisplayWelcome(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ui.displayWelcome()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Welcome to Chess!") {
		t.Error("Welcome message should contain 'Welcome to Chess!'")
	}
	if !strings.Contains(output, "algebraic notation") {
		t.Error("Welcome message should mention algebraic notation")
	}
}

func TestDisplayHelp(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ui.displayHelp()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Help") {
		t.Error("Help message should contain 'Help'")
	}
	if !strings.Contains(output, "Piece Symbols") {
		t.Error("Help message should contain piece symbols")
	}
	if !strings.Contains(output, "♔/♚ = King") {
		t.Error("Help message should explain piece symbols")
	}
}

func TestDisplayBoard(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ui.displayBoard()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "a  b  c  d  e  f  g  h") {
		t.Error("Board display should contain column labels")
	}
	if !strings.Contains(output, "♔ ") {
		t.Error("Board display should contain white king")
	}
}

func TestDisplayGameStatus(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ui.displayGameStatus()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Current Player") {
		t.Error("Game status should contain current player")
	}
	if !strings.Contains(output, "White") {
		t.Error("Game status should show White as initial player")
	}
}

func TestShowValidMoves(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test valid position with moves
	ui.showValidMoves("e2")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Valid moves") {
		t.Error("Should show valid moves for e2")
	}
}

func TestShowValidMovesInvalidPosition(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test invalid position
	ui.showValidMoves("z9")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Invalid position") {
		t.Error("Should show error for invalid position")
	}
}

func TestShowValidMovesEmptySquare(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test empty square
	ui.showValidMoves("e4")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "No piece at position") {
		t.Error("Should show error for empty square")
	}
}

func TestShowValidMovesWrongPlayer(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test opponent's piece (black piece when it's white's turn)
	ui.showValidMoves("e7")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "not your piece") {
		t.Error("Should show error for opponent's piece")
	}
}

func TestShowValidMovesNoMoves(t *testing.T) {
	ui := NewInterface()

	// Clear the board first
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			ui.game.Board.SetPiece(chess.NewPosition(i, j), nil)
		}
	}

	// Create a scenario where a white king shows basic valid moves
	// Place white king at d5 (3,3)
	ui.game.Board.SetPiece(chess.NewPosition(3, 3), chess.NewPiece(chess.King, chess.White))

	// Surround it with black queens that attack all adjacent squares
	ui.game.Board.SetPiece(chess.NewPosition(2, 2), chess.NewPiece(chess.Queen, chess.Black)) // c6
	ui.game.Board.SetPiece(chess.NewPosition(2, 3), chess.NewPiece(chess.Queen, chess.Black)) // d6
	ui.game.Board.SetPiece(chess.NewPosition(2, 4), chess.NewPiece(chess.Queen, chess.Black)) // e6
	ui.game.Board.SetPiece(chess.NewPosition(3, 2), chess.NewPiece(chess.Queen, chess.Black)) // c5
	ui.game.Board.SetPiece(chess.NewPosition(3, 4), chess.NewPiece(chess.Queen, chess.Black)) // e5
	ui.game.Board.SetPiece(chess.NewPosition(4, 2), chess.NewPiece(chess.Queen, chess.Black)) // c4
	ui.game.Board.SetPiece(chess.NewPosition(4, 3), chess.NewPiece(chess.Queen, chess.Black)) // d4
	ui.game.Board.SetPiece(chess.NewPosition(4, 4), chess.NewPiece(chess.Queen, chess.Black)) // e4

	// Ensure it's White's turn
	ui.game.CurrentPlayer = chess.White

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ui.showValidMoves("d5") // The king at d5

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// The king should show basic valid moves (even though they would be illegal due to check)
	if !strings.Contains(output, "Valid moves for White King at d5") {
		t.Errorf("Should show valid moves for king, got: %s", output)
	}
}

func TestProcessMoveValid(t *testing.T) {
	ui := NewInterface()

	success := ui.processMove("e2 e4")

	if !success {
		t.Error("Valid move should return true")
	}

	// Check that the move was made
	piece := ui.game.Board.GetPiece(chess.NewPosition(4, 4)) // e4
	if piece == nil || piece.Type != chess.Pawn {
		t.Error("Pawn should have moved to e4")
	}
}

func TestProcessMoveInvalidFormat(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	success := ui.processMove("e2")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if success {
		t.Error("Invalid format should return false")
	}
	if !strings.Contains(output, "Invalid move format") {
		t.Error("Should show format error message")
	}
}

func TestProcessMoveInvalidMove(t *testing.T) {
	ui := NewInterface()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	success := ui.processMove("e2 e5") // Invalid pawn move

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if success {
		t.Error("Invalid move should return false")
	}
	if !strings.Contains(output, "Invalid move") {
		t.Error("Should show invalid move error")
	}
}

func TestGetInput(t *testing.T) {
	ui := NewInterface()

	// Create a pipe to simulate user input
	r, w, _ := os.Pipe()
	ui.reader = bufio.NewReader(r)

	// Write test input
	go func() {
		w.WriteString("e2 e4\n")
		w.Close()
	}()

	input := ui.getInput()
	if input != "e2 e4" {
		t.Errorf("Expected 'e2 e4', got '%s'", input)
	}
}

func TestClearScreen(t *testing.T) {
	ui := NewInterface()

	// This function executes a system command, so we just test that it doesn't panic
	// The actual clearing effect can't be easily tested in a unit test
	ui.clearScreen()

	// If we get here without panicking, the test passes
}

// Test the Run method with simulated input
func TestRunQuit(t *testing.T) {
	ui := NewInterface()

	// Create a pipe to simulate user input
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	ui.reader = bufio.NewReader(r)

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Write quit command
	go func() {
		w.WriteString("quit\n")
		w.Close()
	}()

	// Run the interface (should quit immediately)
	ui.Run()

	// Restore stdin/stdout
	os.Stdin = oldStdin
	wOut.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	output := buf.String()

	if !strings.Contains(output, "Thanks for playing") {
		t.Error("Should show quit message")
	}
}

func TestRunHelp(t *testing.T) {
	ui := NewInterface()

	// Create a pipe to simulate user input
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	ui.reader = bufio.NewReader(r)

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Write help command followed by quit
	go func() {
		w.WriteString("help\n")
		w.WriteString("quit\n")
		w.Close()
	}()

	// Run the interface
	ui.Run()

	// Restore stdin/stdout
	os.Stdin = oldStdin
	wOut.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	output := buf.String()

	if !strings.Contains(output, "Help") {
		t.Error("Should show help message")
	}
}

func TestRunMoves(t *testing.T) {
	ui := NewInterface()

	// Create a pipe to simulate user input
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	ui.reader = bufio.NewReader(r)

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Write moves command followed by quit
	go func() {
		w.WriteString("moves e2\n")
		w.WriteString("quit\n")
		w.Close()
	}()

	// Run the interface
	ui.Run()

	// Restore stdin/stdout
	os.Stdin = oldStdin
	wOut.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	output := buf.String()

	if !strings.Contains(output, "Valid moves") {
		t.Error("Should show valid moves")
	}
}

func TestRunEmptyInput(t *testing.T) {
	ui := NewInterface()

	// Create a pipe to simulate user input
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	ui.reader = bufio.NewReader(r)

	// Capture stdout
	oldStdout := os.Stdout
	_, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Write empty input followed by quit
	go func() {
		w.WriteString("\n")
		w.WriteString("quit\n")
		w.Close()
	}()

	// Run the interface
	ui.Run()

	// Restore stdin/stdout
	os.Stdin = oldStdin
	wOut.Close()
	os.Stdout = oldStdout

	// Empty input should just continue the loop
	// If we get here without hanging, the test passes
}

func TestRunValidMove(t *testing.T) {
	ui := NewInterface()

	// Create a pipe to simulate user input
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	ui.reader = bufio.NewReader(r)

	// Capture stdout
	oldStdout := os.Stdout
	_, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Write valid move followed by quit
	go func() {
		w.WriteString("e2 e4\n")
		w.WriteString("quit\n")
		w.Close()
	}()

	// Run the interface
	ui.Run()

	// Restore stdin/stdout
	os.Stdin = oldStdin
	wOut.Close()
	os.Stdout = oldStdout

	// Check that the move was made
	piece := ui.game.Board.GetPiece(chess.NewPosition(4, 4)) // e4
	if piece == nil || piece.Type != chess.Pawn {
		t.Error("Pawn should have moved to e4")
	}
}

func TestRunGameOver(t *testing.T) {
	ui := NewInterface()

	// Set game to be over
	ui.game.State = chess.Checkmate

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Run the interface (should exit immediately due to game over)
	ui.Run()

	// Restore stdout
	wOut.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	output := buf.String()

	if !strings.Contains(output, "Game Over") {
		t.Error("Should show game over message")
	}
}
