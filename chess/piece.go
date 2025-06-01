// Package chess implements a complete chess game with all standard rules and AI opponent.
package chess

// Constants for repeated strings
const (
	ColorWhite   = "White"
	ColorBlack   = "Black"
	EmptySquare  = "  "
	UnknownValue = "Unknown"
)

// Color represents the color of a chess piece
type Color int

const (
	// White represents the white player
	White Color = iota
	// Black represents the black player
	Black
)

// String returns the string representation of the color
func (c Color) String() string {
	switch c {
	case White:
		return ColorWhite
	case Black:
		return ColorBlack
	default:
		return UnknownValue
	}
}

// PieceType represents the type of chess piece
type PieceType int

const (
	// King represents the king piece
	King PieceType = iota
	// Queen represents the queen piece
	Queen
	// Rook represents the rook piece
	Rook
	// Bishop represents the bishop piece
	Bishop
	// Knight represents the knight piece
	Knight
	// Pawn represents the pawn piece
	Pawn
)

func (pt PieceType) String() string {
	switch pt {
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	default:
		return "Unknown"
	}
}

// Piece represents a chess piece
type Piece struct {
	Type     PieceType
	Color    Color
	HasMoved bool
}

// NewPiece creates a new chess piece
func NewPiece(pieceType PieceType, color Color) *Piece {
	return &Piece{
		Type:     pieceType,
		Color:    color,
		HasMoved: false,
	}
}

// String returns the string representation of a piece using chess emojis
func (p *Piece) String() string {
	if p == nil {
		return EmptySquare
	}

	switch p.Type {
	case King:
		if p.Color == White {
			return "♔ "
		}
		return "♚ "
	case Queen:
		if p.Color == White {
			return "♕ "
		}
		return "♛ "
	case Rook:
		if p.Color == White {
			return "♖ "
		}
		return "♜ "
	case Bishop:
		if p.Color == White {
			return "♗ "
		}
		return "♝ "
	case Knight:
		if p.Color == White {
			return "♘ "
		}
		return "♞ "
	case Pawn:
		if p.Color == White {
			return "♙ "
		}
		return "♟ "
	default:
		return EmptySquare
	}
}
