package chess

// Color represents the color of a chess piece
type Color int

const (
	White Color = iota
	Black
)

func (c Color) String() string {
	if c == White {
		return "White"
	}
	return "Black"
}

// PieceType represents the type of chess piece
type PieceType int

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
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
		return "  "
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
		return "  "
	}
}
