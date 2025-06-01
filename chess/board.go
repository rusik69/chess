package chess

import (
	"fmt"
	"strings"
)

// Board represents the chess board
type Board struct {
	squares [8][8]*Piece
}

// NewBoard creates a new chess board with pieces in starting positions
func NewBoard() *Board {
	board := &Board{}
	board.setupInitialPosition()
	return board
}

// setupInitialPosition sets up the initial chess position
func (b *Board) setupInitialPosition() {
	// Clear the board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.squares[i][j] = nil
		}
	}

	// Place white pieces
	b.squares[7][0] = NewPiece(Rook, White)
	b.squares[7][1] = NewPiece(Knight, White)
	b.squares[7][2] = NewPiece(Bishop, White)
	b.squares[7][3] = NewPiece(Queen, White)
	b.squares[7][4] = NewPiece(King, White)
	b.squares[7][5] = NewPiece(Bishop, White)
	b.squares[7][6] = NewPiece(Knight, White)
	b.squares[7][7] = NewPiece(Rook, White)

	// Place white pawns
	for i := 0; i < 8; i++ {
		b.squares[6][i] = NewPiece(Pawn, White)
	}

	// Place black pieces
	b.squares[0][0] = NewPiece(Rook, Black)
	b.squares[0][1] = NewPiece(Knight, Black)
	b.squares[0][2] = NewPiece(Bishop, Black)
	b.squares[0][3] = NewPiece(Queen, Black)
	b.squares[0][4] = NewPiece(King, Black)
	b.squares[0][5] = NewPiece(Bishop, Black)
	b.squares[0][6] = NewPiece(Knight, Black)
	b.squares[0][7] = NewPiece(Rook, Black)

	// Place black pawns
	for i := 0; i < 8; i++ {
		b.squares[1][i] = NewPiece(Pawn, Black)
	}
}

// GetPiece returns the piece at the given position
func (b *Board) GetPiece(pos Position) *Piece {
	if !pos.IsValid() {
		return nil
	}
	return b.squares[pos.Row][pos.Col]
}

// SetPiece sets a piece at the given position
func (b *Board) SetPiece(pos Position, piece *Piece) {
	if pos.IsValid() {
		b.squares[pos.Row][pos.Col] = piece
	}
}

// MovePiece moves a piece from one position to another
func (b *Board) MovePiece(from, to Position) bool {
	if !from.IsValid() || !to.IsValid() {
		return false
	}

	piece := b.GetPiece(from)
	if piece == nil {
		return false
	}

	b.SetPiece(to, piece)
	b.SetPiece(from, nil)
	piece.HasMoved = true

	return true
}

// String returns a string representation of the board
func (b *Board) String() string {
	var sb strings.Builder

	sb.WriteString("   a  b  c  d  e  f  g  h\n")
	sb.WriteString("  ┌──┬──┬──┬──┬──┬──┬──┬──┐\n")

	for row := 0; row < 8; row++ {
		sb.WriteString(fmt.Sprintf("%d │", 8-row))

		for col := 0; col < 8; col++ {
			piece := b.squares[row][col]
			if piece == nil {
				sb.WriteString("  ")
			} else {
				sb.WriteString(piece.String())
			}
			sb.WriteString("│")
		}

		sb.WriteString(fmt.Sprintf(" %d\n", 8-row))

		if row < 7 {
			sb.WriteString("  ├──┼──┼──┼──┼──┼──┼──┼──┤\n")
		}
	}

	sb.WriteString("  └──┴──┴──┴──┴──┴──┴──┴──┘\n")
	sb.WriteString("   a  b  c  d  e  f  g  h\n")

	return sb.String()
}
