package chess

import (
	"fmt"
	"strings"
)

// GameState represents the current state of the game
type GameState int

const (
	Playing GameState = iota
	Check
	Checkmate
	Stalemate
	Draw
)

func (gs GameState) String() string {
	switch gs {
	case Playing:
		return "Playing"
	case Check:
		return "Check"
	case Checkmate:
		return "Checkmate"
	case Stalemate:
		return "Stalemate"
	case Draw:
		return "Draw"
	default:
		return "Unknown"
	}
}

// Game represents a chess game
type Game struct {
	Board         *Board
	CurrentPlayer Color
	State         GameState
	MoveHistory   []Move
}

// NewGame creates a new chess game
func NewGame() *Game {
	return &Game{
		Board:         NewBoard(),
		CurrentPlayer: White,
		State:         Playing,
		MoveHistory:   make([]Move, 0),
	}
}

// MakeMove attempts to make a move and returns whether it was successful
func (g *Game) MakeMove(from, to string) error {
	fromPos, err := FromAlgebraic(from)
	if err != nil {
		return fmt.Errorf("invalid from position: %v", err)
	}

	toPos, err := FromAlgebraic(to)
	if err != nil {
		return fmt.Errorf("invalid to position: %v", err)
	}

	move := NewMove(fromPos, toPos)

	if !g.Board.IsValidMove(move, g.CurrentPlayer) {
		return fmt.Errorf("invalid move")
	}

	// Make the move
	g.Board.MovePiece(move.From, move.To)
	g.MoveHistory = append(g.MoveHistory, move)

	// Switch players
	if g.CurrentPlayer == White {
		g.CurrentPlayer = Black
	} else {
		g.CurrentPlayer = White
	}

	// Update game state
	g.updateGameState()

	return nil
}

// updateGameState updates the current game state
func (g *Game) updateGameState() {
	if g.isInCheck(g.CurrentPlayer) {
		if g.hasValidMoves(g.CurrentPlayer) {
			g.State = Check
		} else {
			g.State = Checkmate
		}
	} else if !g.hasValidMoves(g.CurrentPlayer) {
		g.State = Stalemate
	} else {
		g.State = Playing
	}
}

// isInCheck checks if the given player's king is in check
func (g *Game) isInCheck(player Color) bool {
	// Find the king
	var kingPos Position
	found := false
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			pos := NewPosition(row, col)
			piece := g.Board.GetPiece(pos)
			if piece != nil && piece.Type == King && piece.Color == player {
				kingPos = pos
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return false
	}

	// Check if any opponent piece can attack the king
	opponent := Black
	if player == Black {
		opponent = White
	}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			pos := NewPosition(row, col)
			piece := g.Board.GetPiece(pos)
			if piece != nil && piece.Color == opponent {
				move := NewMove(pos, kingPos)
				if g.Board.IsValidMove(move, opponent) {
					return true
				}
			}
		}
	}

	return false
}

// hasValidMoves checks if the player has any valid moves
func (g *Game) hasValidMoves(player Color) bool {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			pos := NewPosition(row, col)
			piece := g.Board.GetPiece(pos)
			if piece != nil && piece.Color == player {
				moves := g.Board.GetValidMoves(pos)
				for _, move := range moves {
					// Try the move and see if it leaves the king in check
					originalPiece := g.Board.GetPiece(move.To)
					g.Board.MovePiece(move.From, move.To)

					inCheck := g.isInCheck(player)

					// Undo the move
					g.Board.MovePiece(move.To, move.From)
					g.Board.SetPiece(move.To, originalPiece)

					if !inCheck {
						return true
					}
				}
			}
		}
	}
	return false
}

// GetGameStatus returns a string describing the current game status
func (g *Game) GetGameStatus() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Current Player: %s\n", g.CurrentPlayer))
	sb.WriteString(fmt.Sprintf("Game State: %s\n", g.State))

	if g.State == Check {
		sb.WriteString(fmt.Sprintf("%s is in check!\n", g.CurrentPlayer))
	} else if g.State == Checkmate {
		winner := Black
		if g.CurrentPlayer == Black {
			winner = White
		}
		sb.WriteString(fmt.Sprintf("Checkmate! %s wins!\n", winner))
	} else if g.State == Stalemate {
		sb.WriteString("Stalemate! The game is a draw.\n")
	}

	return sb.String()
}

// IsGameOver returns true if the game is over
func (g *Game) IsGameOver() bool {
	return g.State == Checkmate || g.State == Stalemate || g.State == Draw
}
