package chess

import (
	"strings"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	if game.CurrentPlayer != White {
		t.Errorf("Expected current player to be White, got %v", game.CurrentPlayer)
	}

	if game.State != Playing {
		t.Errorf("Expected game state to be Playing, got %v", game.State)
	}

	// Test initial board setup
	// Check white king position
	whiteKing := game.Board.GetPiece(NewPosition(7, 4))
	if whiteKing == nil || whiteKing.Type != King || whiteKing.Color != White {
		t.Error("White king not in correct starting position")
	}

	// Check black king position
	blackKing := game.Board.GetPiece(NewPosition(0, 4))
	if blackKing == nil || blackKing.Type != King || blackKing.Color != Black {
		t.Error("Black king not in correct starting position")
	}
}

func TestValidPawnMove(t *testing.T) {
	game := NewGame()

	// Test valid pawn move
	err := game.MakeMove("e2", "e4")
	if err != nil {
		t.Errorf("Expected valid pawn move, got error: %v", err)
	}

	// Check that the pawn moved
	piece := game.Board.GetPiece(NewPosition(4, 4)) // e4
	if piece == nil || piece.Type != Pawn || piece.Color != White {
		t.Error("Pawn did not move to correct position")
	}

	// Check that current player switched
	if game.CurrentPlayer != Black {
		t.Errorf("Expected current player to be Black after move, got %v", game.CurrentPlayer)
	}
}

func TestInvalidMove(t *testing.T) {
	game := NewGame()

	// Test invalid move (trying to move opponent's piece)
	err := game.MakeMove("e7", "e5")
	if err == nil {
		t.Error("Expected error when trying to move opponent's piece")
	}

	// Test invalid move (illegal pawn move)
	err = game.MakeMove("e2", "e5")
	if err == nil {
		t.Error("Expected error for illegal pawn move")
	}
}

func TestPositionConversion(t *testing.T) {
	// Test algebraic notation conversion
	pos, err := FromAlgebraic("e4")
	if err != nil {
		t.Errorf("Error converting algebraic notation: %v", err)
	}

	expected := NewPosition(4, 4)
	if pos != expected {
		t.Errorf("Expected position %v, got %v", expected, pos)
	}

	// Test position to string
	str := pos.String()
	if str != "e4" {
		t.Errorf("Expected string 'e4', got '%s'", str)
	}
}

// Test String methods for Color and PieceType
func TestColorString(t *testing.T) {
	if White.String() != "White" {
		t.Errorf("Expected 'White', got '%s'", White.String())
	}
	if Black.String() != "Black" {
		t.Errorf("Expected 'Black', got '%s'", Black.String())
	}
}

func TestPieceTypeString(t *testing.T) {
	tests := []struct {
		pieceType PieceType
		expected  string
	}{
		{King, "King"},
		{Queen, "Queen"},
		{Rook, "Rook"},
		{Bishop, "Bishop"},
		{Knight, "Knight"},
		{Pawn, "Pawn"},
		{PieceType(99), "Unknown"}, // Invalid piece type
	}

	for _, test := range tests {
		if test.pieceType.String() != test.expected {
			t.Errorf("Expected '%s', got '%s'", test.expected, test.pieceType.String())
		}
	}
}

func TestGameStateString(t *testing.T) {
	tests := []struct {
		state    GameState
		expected string
	}{
		{Playing, "Playing"},
		{Check, "Check"},
		{Checkmate, "Checkmate"},
		{Stalemate, "Stalemate"},
		{Draw, "Draw"},
		{GameState(99), "Unknown"}, // Invalid game state
	}

	for _, test := range tests {
		if test.state.String() != test.expected {
			t.Errorf("Expected '%s', got '%s'", test.expected, test.state.String())
		}
	}
}

func TestPieceString(t *testing.T) {
	// Test nil piece
	var nilPiece *Piece
	if nilPiece.String() != "  " {
		t.Errorf("Expected '  ' for nil piece, got '%s'", nilPiece.String())
	}

	// Test all piece types and colors
	tests := []struct {
		piece    *Piece
		expected string
	}{
		{NewPiece(King, White), "♔ "},
		{NewPiece(Queen, White), "♕ "},
		{NewPiece(Rook, White), "♖ "},
		{NewPiece(Bishop, White), "♗ "},
		{NewPiece(Knight, White), "♘ "},
		{NewPiece(Pawn, White), "♙ "},
		{NewPiece(King, Black), "♚ "},
		{NewPiece(Queen, Black), "♛ "},
		{NewPiece(Rook, Black), "♜ "},
		{NewPiece(Bishop, Black), "♝ "},
		{NewPiece(Knight, Black), "♞ "},
		{NewPiece(Pawn, Black), "♟ "},
	}

	for _, test := range tests {
		if test.piece.String() != test.expected {
			t.Errorf("Expected '%s', got '%s'", test.expected, test.piece.String())
		}
	}
}

func TestBoardString(t *testing.T) {
	board := NewBoard()
	boardStr := board.String()

	// Check that board string contains expected elements
	if !strings.Contains(boardStr, "a  b  c  d  e  f  g  h") {
		t.Error("Board string should contain column labels")
	}
	if !strings.Contains(boardStr, "♔ ") {
		t.Error("Board string should contain white king")
	}
	if !strings.Contains(boardStr, "♚ ") {
		t.Error("Board string should contain black king")
	}
}

func TestPositionInvalid(t *testing.T) {
	// Test invalid position string
	invalidPos := NewPosition(-1, -1)
	if invalidPos.String() != "invalid" {
		t.Errorf("Expected 'invalid', got '%s'", invalidPos.String())
	}

	// Test FromAlgebraic with invalid inputs
	_, err := FromAlgebraic("z9")
	if err == nil {
		t.Error("Expected error for invalid algebraic notation")
	}

	_, err = FromAlgebraic("abc")
	if err == nil {
		t.Error("Expected error for invalid notation length")
	}
}

func TestBoardGetPieceInvalid(t *testing.T) {
	board := NewBoard()

	// Test getting piece from invalid position
	piece := board.GetPiece(NewPosition(-1, -1))
	if piece != nil {
		t.Error("Expected nil for invalid position")
	}
}

func TestBoardMovePieceInvalid(t *testing.T) {
	board := NewBoard()

	// Test moving from invalid position
	success := board.MovePiece(NewPosition(-1, -1), NewPosition(0, 0))
	if success {
		t.Error("Expected false for invalid from position")
	}

	// Test moving to invalid position
	success = board.MovePiece(NewPosition(0, 0), NewPosition(-1, -1))
	if success {
		t.Error("Expected false for invalid to position")
	}

	// Test moving from empty square
	success = board.MovePiece(NewPosition(3, 3), NewPosition(4, 4))
	if success {
		t.Error("Expected false for moving from empty square")
	}
}

func TestMakeMoveInvalidPositions(t *testing.T) {
	game := NewGame()

	// Test invalid from position
	err := game.MakeMove("z9", "e4")
	if err == nil {
		t.Error("Expected error for invalid from position")
	}

	// Test invalid to position
	err = game.MakeMove("e2", "z9")
	if err == nil {
		t.Error("Expected error for invalid to position")
	}
}

func TestPawnMoves(t *testing.T) {
	game := NewGame()

	// Test pawn two-square move
	err := game.MakeMove("e2", "e4")
	if err != nil {
		t.Errorf("Expected valid pawn two-square move, got error: %v", err)
	}

	// Test pawn one-square move
	err = game.MakeMove("e7", "e6")
	if err != nil {
		t.Errorf("Expected valid pawn one-square move, got error: %v", err)
	}

	// Test pawn diagonal capture
	game.MakeMove("d2", "d4")
	game.MakeMove("f7", "f5")
	err = game.MakeMove("e4", "f5") // Capture
	if err != nil {
		t.Errorf("Expected valid pawn capture, got error: %v", err)
	}
}

func TestKnightMoves(t *testing.T) {
	game := NewGame()

	// Test knight move
	err := game.MakeMove("g1", "f3")
	if err != nil {
		t.Errorf("Expected valid knight move, got error: %v", err)
	}

	err = game.MakeMove("b8", "c6")
	if err != nil {
		t.Errorf("Expected valid knight move, got error: %v", err)
	}
}

func TestBishopMoves(t *testing.T) {
	game := NewGame()

	// Move pawns to open diagonal
	game.MakeMove("e2", "e4")
	game.MakeMove("e7", "e5")

	// Test bishop move
	err := game.MakeMove("f1", "c4")
	if err != nil {
		t.Errorf("Expected valid bishop move, got error: %v", err)
	}
}

func TestRookMoves(t *testing.T) {
	game := NewGame()

	// Move pawns to open file
	game.MakeMove("a2", "a4")
	game.MakeMove("a7", "a5")

	// Test rook move
	err := game.MakeMove("a1", "a3")
	if err != nil {
		t.Errorf("Expected valid rook move, got error: %v", err)
	}
}

func TestQueenMoves(t *testing.T) {
	game := NewGame()

	// Move pawns to open diagonal and file
	game.MakeMove("d2", "d4")
	game.MakeMove("d7", "d5")

	// Test queen move
	err := game.MakeMove("d1", "d3")
	if err != nil {
		t.Errorf("Expected valid queen move, got error: %v", err)
	}
}

func TestQueenMovesAdvanced(t *testing.T) {
	board := NewBoard()

	// Clear some pawns to test queen movement
	board.SetPiece(NewPosition(6, 3), nil) // Remove d2 pawn

	// Test queen moving like a rook
	move := NewMove(NewPosition(7, 3), NewPosition(5, 3)) // d1 to d3
	if !board.IsValidMove(move, White) {
		t.Error("Queen should be able to move like a rook")
	}
}

func TestKingMoves(t *testing.T) {
	game := NewGame()

	// Move pawns to allow king movement
	game.MakeMove("e2", "e4")
	game.MakeMove("e7", "e5")

	// Test king move
	err := game.MakeMove("e1", "e2")
	if err != nil {
		t.Errorf("Expected valid king move, got error: %v", err)
	}
}

func TestPathBlocked(t *testing.T) {
	game := NewGame()

	// Try to move rook through pawn (should fail)
	err := game.MakeMove("a1", "a4")
	if err == nil {
		t.Error("Expected error for blocked path")
	}

	// Try to move bishop through pawn (should fail)
	err = game.MakeMove("c1", "f4")
	if err == nil {
		t.Error("Expected error for blocked path")
	}
}

func TestGameStatus(t *testing.T) {
	game := NewGame()

	// Test initial game status
	status := game.GetGameStatus()
	if !strings.Contains(status, "White") {
		t.Error("Game status should contain current player")
	}
	if !strings.Contains(status, "Playing") {
		t.Error("Game status should contain game state")
	}
}

func TestIsGameOver(t *testing.T) {
	game := NewGame()

	// Initially game should not be over
	if game.IsGameOver() {
		t.Error("Game should not be over initially")
	}

	// Test with different game states
	game.State = Checkmate
	if !game.IsGameOver() {
		t.Error("Game should be over in checkmate")
	}

	game.State = Stalemate
	if !game.IsGameOver() {
		t.Error("Game should be over in stalemate")
	}

	game.State = Draw
	if !game.IsGameOver() {
		t.Error("Game should be over in draw")
	}
}

func TestCheckDetection(t *testing.T) {
	// Create a custom board position with king in check
	game := NewGame()

	// Clear the board and set up a check scenario
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			game.Board.SetPiece(NewPosition(i, j), nil)
		}
	}

	// Place white king and black queen to create check
	game.Board.SetPiece(NewPosition(4, 4), NewPiece(King, White))
	game.Board.SetPiece(NewPosition(4, 0), NewPiece(Queen, Black))

	// Check if white king is in check
	if !game.isInCheck(White) {
		t.Error("White king should be in check")
	}
}

func TestGetValidMoves(t *testing.T) {
	game := NewGame()

	// Test getting valid moves for a pawn
	moves := game.Board.GetValidMoves(NewPosition(6, 4)) // e2 pawn
	if len(moves) == 0 {
		t.Error("Pawn should have valid moves")
	}

	// Test getting valid moves for empty square
	moves = game.Board.GetValidMoves(NewPosition(4, 4)) // empty square
	if len(moves) != 0 {
		t.Error("Empty square should have no valid moves")
	}
}

func TestInvalidMoveScenarios(t *testing.T) {
	game := NewGame()

	// Test moving to square occupied by own piece
	err := game.MakeMove("b1", "c3") // Knight move
	if err != nil {
		t.Errorf("Expected valid knight move, got error: %v", err)
	}

	// Try to move knight to square occupied by own piece
	err = game.MakeMove("g8", "f6") // Black knight
	if err != nil {
		t.Errorf("Expected valid knight move, got error: %v", err)
	}

	// Try invalid knight move
	err = game.MakeMove("c3", "c4") // Not a valid knight move
	if err == nil {
		t.Error("Expected error for invalid knight move")
	}
}

func TestBishopPathClear(t *testing.T) {
	board := NewBoard()

	// Test diagonal path clear
	from := NewPosition(7, 2) // c1 bishop
	to := NewPosition(4, 5)   // f4

	// Path should be blocked by pawn
	if board.isPathClear(from, to) {
		t.Error("Path should be blocked by pawn")
	}

	// Clear the path
	board.SetPiece(NewPosition(6, 3), nil) // Remove d2 pawn
	board.SetPiece(NewPosition(5, 4), nil) // Remove e3 if there was a piece

	// Now path should be clear
	if !board.isPathClear(from, to) {
		t.Error("Path should be clear now")
	}
}

func TestGameStatusMessages(t *testing.T) {
	game := NewGame()

	// Test check status
	game.State = Check
	status := game.GetGameStatus()
	if !strings.Contains(status, "in check") {
		t.Error("Should show check message")
	}

	// Test checkmate status
	game.State = Checkmate
	game.CurrentPlayer = Black // Black is in checkmate, so White wins
	status = game.GetGameStatus()
	if !strings.Contains(status, "Checkmate") {
		t.Error("Should show checkmate message")
	}
	if !strings.Contains(status, "White wins") {
		t.Error("Should show winner")
	}

	// Test stalemate status
	game.State = Stalemate
	status = game.GetGameStatus()
	if !strings.Contains(status, "Stalemate") {
		t.Error("Should show stalemate message")
	}
	if !strings.Contains(status, "draw") {
		t.Error("Should mention draw")
	}
}

func TestUpdateGameState(t *testing.T) {
	game := NewGame()

	// Test that updateGameState doesn't crash
	game.updateGameState()

	// Initially should be in Playing state
	if game.State != Playing {
		t.Error("Initial state should be Playing")
	}
}

func TestHasValidMovesWithCheck(t *testing.T) {
	game := NewGame()

	// Clear the board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			game.Board.SetPiece(NewPosition(i, j), nil)
		}
	}

	// Set up a position where king is in check but can move
	game.Board.SetPiece(NewPosition(4, 4), NewPiece(King, White))
	game.Board.SetPiece(NewPosition(4, 0), NewPiece(Queen, Black))

	// King should have valid moves (can move to escape check)
	if !game.hasValidMoves(White) {
		t.Error("King should have valid moves to escape check")
	}
}

func TestIsInCheckNoKing(t *testing.T) {
	game := NewGame()

	// Clear the board completely
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			game.Board.SetPiece(NewPosition(i, j), nil)
		}
	}

	// No king on board - should return false
	if game.isInCheck(White) {
		t.Error("Should return false when no king on board")
	}
}

func TestPawnMovesEdgeCases(t *testing.T) {
	board := NewBoard()

	// Test pawn trying to move backwards (should fail)
	move := NewMove(NewPosition(6, 4), NewPosition(7, 4)) // e2 to e1
	if board.IsValidMove(move, White) {
		t.Error("Pawn should not be able to move backwards")
	}

	// Test pawn trying to capture forward (should fail)
	move = NewMove(NewPosition(6, 4), NewPosition(5, 4)) // e2 to e3 with piece there
	board.SetPiece(NewPosition(5, 4), NewPiece(Pawn, Black))
	if board.IsValidMove(move, White) {
		t.Error("Pawn should not be able to capture forward")
	}

	// Test pawn trying to move two squares when not on starting row
	board.MovePiece(NewPosition(6, 4), NewPosition(5, 4)) // Move pawn to e3
	board.SetPiece(NewPosition(5, 4), nil)                // Clear the square
	board.SetPiece(NewPosition(5, 4), NewPiece(Pawn, White))
	move = NewMove(NewPosition(5, 4), NewPosition(3, 4)) // e3 to e5 (two squares)
	if board.IsValidMove(move, White) {
		t.Error("Pawn should not be able to move two squares from non-starting position")
	}
}

func TestBishopInvalidMoves(t *testing.T) {
	board := NewBoard()

	// Test bishop trying to move like a rook (should fail)
	// First clear the path
	board.SetPiece(NewPosition(6, 2), nil)                // Remove c2 pawn
	move := NewMove(NewPosition(7, 2), NewPosition(4, 2)) // c1 to c5 (vertical)
	if board.IsValidMove(move, White) {
		t.Error("Bishop should not be able to move vertically")
	}

	// Test bishop trying to move like a knight (should fail)
	move = NewMove(NewPosition(7, 2), NewPosition(5, 3)) // c1 to d6 (knight move)
	if board.IsValidMove(move, White) {
		t.Error("Bishop should not be able to move like a knight")
	}
}

func TestRookInvalidMoves(t *testing.T) {
	board := NewBoard()

	// Test rook trying to move diagonally (should fail)
	// First clear the path
	board.SetPiece(NewPosition(6, 0), nil)                // Remove a2 pawn
	move := NewMove(NewPosition(7, 0), NewPosition(5, 2)) // a1 to c3 (diagonal)
	if board.IsValidMove(move, White) {
		t.Error("Rook should not be able to move diagonally")
	}
}

func TestKingInvalidMoves(t *testing.T) {
	board := NewBoard()

	// Test king trying to move more than one square (should fail)
	// First clear the path
	board.SetPiece(NewPosition(6, 4), nil)                // Remove e2 pawn
	move := NewMove(NewPosition(7, 4), NewPosition(5, 4)) // e1 to e3 (two squares)
	if board.IsValidMove(move, White) {
		t.Error("King should not be able to move more than one square")
	}
}

func TestPathClearEdgeCases(t *testing.T) {
	board := NewBoard()

	// Test path clear for same position (should be true)
	pos := NewPosition(4, 4)
	if !board.isPathClear(pos, pos) {
		t.Error("Path should be clear for same position")
	}

	// Test path clear for adjacent positions (should be true)
	from := NewPosition(4, 4)
	to := NewPosition(4, 5)
	if !board.isPathClear(from, to) {
		t.Error("Path should be clear for adjacent positions")
	}
}

func TestMoveHistoryTracking(t *testing.T) {
	game := NewGame()

	// Make a move and check history
	initialHistoryLength := len(game.MoveHistory)
	err := game.MakeMove("e2", "e4")
	if err != nil {
		t.Errorf("Expected valid move, got error: %v", err)
	}

	if len(game.MoveHistory) != initialHistoryLength+1 {
		t.Error("Move history should be updated")
	}

	lastMove := game.MoveHistory[len(game.MoveHistory)-1]
	if lastMove.From.String() != "e2" || lastMove.To.String() != "e4" {
		t.Error("Move history should contain correct move")
	}
}

func TestPieceHasMovedFlag(t *testing.T) {
	board := NewBoard()

	// Get a piece that hasn't moved
	piece := board.GetPiece(NewPosition(6, 4)) // e2 pawn
	if piece.HasMoved {
		t.Error("Piece should not have moved initially")
	}

	// Move the piece
	board.MovePiece(NewPosition(6, 4), NewPosition(4, 4))

	// Check that HasMoved flag is set
	if !piece.HasMoved {
		t.Error("Piece should have HasMoved flag set after moving")
	}
}

func TestInvalidPieceType(t *testing.T) {
	board := NewBoard()

	// Create a piece with invalid type and test move validation
	invalidPiece := &Piece{Type: PieceType(99), Color: White, HasMoved: false}
	board.SetPiece(NewPosition(4, 4), invalidPiece)

	move := NewMove(NewPosition(4, 4), NewPosition(4, 5))
	if board.IsValidMove(move, White) {
		t.Error("Invalid piece type should not have valid moves")
	}
}

func TestSetPieceInvalidPosition(t *testing.T) {
	board := NewBoard()

	// Try to set piece at invalid position
	piece := NewPiece(Pawn, White)
	board.SetPiece(NewPosition(-1, -1), piece)

	// Should not crash, and piece should not be placed
	retrievedPiece := board.GetPiece(NewPosition(-1, -1))
	if retrievedPiece != nil {
		t.Error("Should not be able to set piece at invalid position")
	}
}

func TestComplexGameScenario(t *testing.T) {
	game := NewGame()

	// Play a series of moves to test game flow
	moves := [][]string{
		{"e2", "e4"}, // 1. e4
		{"e7", "e5"}, // 1... e5
		{"g1", "f3"}, // 2. Nf3
		{"b8", "c6"}, // 2... Nc6
		{"f1", "b5"}, // 3. Bb5
		{"a7", "a6"}, // 3... a6
		{"b5", "a4"}, // 4. Ba4
		{"g8", "f6"}, // 4... Nf6
	}

	for i, move := range moves {
		err := game.MakeMove(move[0], move[1])
		if err != nil {
			t.Errorf("Move %d (%s %s) failed: %v", i+1, move[0], move[1], err)
		}
	}

	// Verify some pieces are in expected positions
	// White bishop should be on a4
	piece := game.Board.GetPiece(NewPosition(4, 0)) // a4
	if piece == nil || piece.Type != Bishop || piece.Color != White {
		t.Error("White bishop should be on a4")
	}

	// Black knight should be on f6
	piece = game.Board.GetPiece(NewPosition(2, 5)) // f6
	if piece == nil || piece.Type != Knight || piece.Color != Black {
		t.Error("Black knight should be on f6")
	}
}
