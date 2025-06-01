package chess

import "fmt"

// Position represents a position on the chess board
type Position struct {
	Row int
	Col int
}

// NewPosition creates a new position
func NewPosition(row, col int) Position {
	return Position{Row: row, Col: col}
}

// IsValid checks if the position is within the board boundaries
func (p Position) IsValid() bool {
	return p.Row >= 0 && p.Row < 8 && p.Col >= 0 && p.Col < 8
}

// String returns the algebraic notation of the position (e.g., "e4")
func (p Position) String() string {
	if !p.IsValid() {
		return "invalid"
	}
	col := string(rune('a' + p.Col))
	row := fmt.Sprintf("%d", 8-p.Row)
	return col + row
}

// FromAlgebraic converts algebraic notation to Position
func FromAlgebraic(notation string) (Position, error) {
	if len(notation) != 2 {
		return Position{}, fmt.Errorf("invalid notation: %s", notation)
	}

	col := int(notation[0] - 'a')
	row := 8 - int(notation[1]-'0')

	pos := Position{Row: row, Col: col}
	if !pos.IsValid() {
		return Position{}, fmt.Errorf("invalid position: %s", notation)
	}

	return pos, nil
}
