package chess

import "math"

// Move represents a chess move
type Move struct {
	From Position
	To   Position
}

// NewMove creates a new move
func NewMove(from, to Position) Move {
	return Move{From: from, To: to}
}

// IsValidMove checks if a move is valid for the given board state
func (b *Board) IsValidMove(move Move, currentPlayer Color) bool {
	piece := b.GetPiece(move.From)
	if piece == nil {
		return false
	}

	// Check if it's the current player's piece
	if piece.Color != currentPlayer {
		return false
	}

	// Check if destination has own piece
	destPiece := b.GetPiece(move.To)
	if destPiece != nil && destPiece.Color == piece.Color {
		return false
	}

	// Check piece-specific movement rules
	switch piece.Type {
	case Pawn:
		return b.isValidPawnMove(move, piece)
	case Rook:
		return b.isValidRookMove(move)
	case Knight:
		return b.isValidKnightMove(move)
	case Bishop:
		return b.isValidBishopMove(move)
	case Queen:
		return b.isValidQueenMove(move)
	case King:
		return b.isValidKingMove(move)
	}

	return false
}

// isValidPawnMove validates pawn movement
func (b *Board) isValidPawnMove(move Move, piece *Piece) bool {
	rowDiff := move.To.Row - move.From.Row
	colDiff := move.To.Col - move.From.Col

	direction := -1 // White moves up (decreasing row numbers)
	startRow := 6   // White pawns start at row 6
	if piece.Color == Black {
		direction = 1 // Black moves down (increasing row numbers)
		startRow = 1  // Black pawns start at row 1
	}

	// Forward move
	if colDiff == 0 {
		// One square forward
		if rowDiff == direction {
			return b.GetPiece(move.To) == nil
		}
		// Two squares forward from starting position
		if rowDiff == 2*direction && move.From.Row == startRow {
			return b.GetPiece(move.To) == nil && b.GetPiece(NewPosition(move.From.Row+direction, move.From.Col)) == nil
		}
	}

	// Diagonal capture
	if math.Abs(float64(colDiff)) == 1 && rowDiff == direction {
		destPiece := b.GetPiece(move.To)
		return destPiece != nil && destPiece.Color != piece.Color
	}

	return false
}

// isValidRookMove validates rook movement
func (b *Board) isValidRookMove(move Move) bool {
	// Rook moves horizontally or vertically
	if move.From.Row != move.To.Row && move.From.Col != move.To.Col {
		return false
	}

	return b.isPathClear(move.From, move.To)
}

// isValidKnightMove validates knight movement
func (b *Board) isValidKnightMove(move Move) bool {
	rowDiff := int(math.Abs(float64(move.To.Row - move.From.Row)))
	colDiff := int(math.Abs(float64(move.To.Col - move.From.Col)))

	return (rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2)
}

// isValidBishopMove validates bishop movement
func (b *Board) isValidBishopMove(move Move) bool {
	rowDiff := int(math.Abs(float64(move.To.Row - move.From.Row)))
	colDiff := int(math.Abs(float64(move.To.Col - move.From.Col)))

	// Bishop moves diagonally
	if rowDiff != colDiff {
		return false
	}

	return b.isPathClear(move.From, move.To)
}

// isValidQueenMove validates queen movement
func (b *Board) isValidQueenMove(move Move) bool {
	// Queen moves like rook or bishop
	return b.isValidRookMove(move) || b.isValidBishopMove(move)
}

// isValidKingMove validates king movement
func (b *Board) isValidKingMove(move Move) bool {
	rowDiff := int(math.Abs(float64(move.To.Row - move.From.Row)))
	colDiff := int(math.Abs(float64(move.To.Col - move.From.Col)))

	// King moves one square in any direction
	return rowDiff <= 1 && colDiff <= 1
}

// isPathClear checks if the path between two positions is clear
func (b *Board) isPathClear(from, to Position) bool {
	rowStep := 0
	colStep := 0

	if to.Row > from.Row {
		rowStep = 1
	} else if to.Row < from.Row {
		rowStep = -1
	}

	if to.Col > from.Col {
		colStep = 1
	} else if to.Col < from.Col {
		colStep = -1
	}

	currentRow := from.Row + rowStep
	currentCol := from.Col + colStep

	for currentRow != to.Row || currentCol != to.Col {
		if b.GetPiece(NewPosition(currentRow, currentCol)) != nil {
			return false
		}
		currentRow += rowStep
		currentCol += colStep
	}

	return true
}

// GetValidMoves returns all valid moves for a piece at the given position
func (b *Board) GetValidMoves(pos Position) []Move {
	var moves []Move
	piece := b.GetPiece(pos)
	if piece == nil {
		return moves
	}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			to := NewPosition(row, col)
			move := NewMove(pos, to)
			if b.IsValidMove(move, piece.Color) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}
