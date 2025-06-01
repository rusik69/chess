package chess

import (
	"math"
	"sort"
)

// TranspositionEntry represents an entry in the transposition table
type TranspositionEntry struct {
	depth int
	score float64
	flag  int // 0 = exact, 1 = lower bound, 2 = upper bound
}

// AI represents the computer player
type AI struct {
	color              Color
	depth              int
	transpositionTable map[uint64]TranspositionEntry
	killerMoves        [10][2]Move // killer moves for each depth
	historyTable       map[Move]int
}

// NewAI creates a new AI player
func NewAI(color Color, depth int) *AI {
	return &AI{
		color:              color,
		depth:              depth,
		transpositionTable: make(map[uint64]TranspositionEntry),
		historyTable:       make(map[Move]int),
	}
}

// GetBestMove returns the best move for the AI player using iterative deepening
func (ai *AI) GetBestMove(game *Game) (Position, Position, bool) {
	if game.CurrentPlayer != ai.color {
		return Position{}, Position{}, false
	}

	allMoves := ai.getAllPossibleMoves(game)
	if len(allMoves) == 0 {
		return Position{}, Position{}, false
	}

	// Clear killer moves for new search
	ai.killerMoves = [10][2]Move{}

	var bestMove Move
	bestScore := math.Inf(-1)

	// Iterative deepening - start with depth 1 and increase
	for currentDepth := 1; currentDepth <= ai.depth; currentDepth++ {
		tempBestMove := Move{}
		tempBestScore := math.Inf(-1)

		// Order moves for better alpha-beta pruning
		orderedMoves := ai.orderMoves(allMoves, game, 0)

		for _, move := range orderedMoves {
			gameCopy := ai.copyGame(game)
			err := gameCopy.MakeMove(move.From.String(), move.To.String())
			if err != nil {
				continue
			}

			score := ai.minimax(gameCopy, currentDepth-1, false, math.Inf(-1), math.Inf(1), 1)

			if score > tempBestScore {
				tempBestScore = score
				tempBestMove = move
			}
		}

		// Update best move if we found a better one
		if tempBestScore > bestScore {
			bestScore = tempBestScore
			bestMove = tempBestMove
		}
	}

	return bestMove.From, bestMove.To, true
}

// minimax implements the minimax algorithm with alpha-beta pruning and optimizations
func (ai *AI) minimax(game *Game, depth int, isMaximizing bool, alpha, beta float64, ply int) float64 {
	// Check transposition table
	hash := ai.hashPosition(game)
	if entry, exists := ai.transpositionTable[hash]; exists && entry.depth >= depth {
		switch entry.flag {
		case 0: // exact
			return entry.score
		case 1: // lower bound
			alpha = math.Max(alpha, entry.score)
		case 2: // upper bound
			beta = math.Min(beta, entry.score)
		}
		if alpha >= beta {
			return entry.score
		}
	}

	if depth == 0 || game.State != Playing {
		score := ai.evaluatePosition(game)
		// Store in transposition table
		ai.transpositionTable[hash] = TranspositionEntry{
			depth: depth,
			score: score,
			flag:  0, // exact
		}
		return score
	}

	moves := ai.getAllPossibleMoves(game)
	if len(moves) == 0 {
		if ai.isInCheck(game, game.CurrentPlayer) {
			// Checkmate - prefer shorter mates
			if isMaximizing {
				return -1000.0 + float64(ply)
			}
			return 1000.0 - float64(ply)
		}
		return 0.0 // Stalemate
	}

	// Order moves for better pruning
	orderedMoves := ai.orderMoves(moves, game, ply)

	originalAlpha := alpha
	bestScore := math.Inf(-1)
	if !isMaximizing {
		bestScore = math.Inf(1)
	}

	for _, move := range orderedMoves {
		gameCopy := ai.copyGame(game)
		err := gameCopy.MakeMove(move.From.String(), move.To.String())
		if err != nil {
			continue
		}

		var score float64
		if isMaximizing {
			score = ai.minimax(gameCopy, depth-1, false, alpha, beta, ply+1)
			if score > bestScore {
				bestScore = score
			}
			alpha = math.Max(alpha, score)
		} else {
			score = ai.minimax(gameCopy, depth-1, true, alpha, beta, ply+1)
			if score < bestScore {
				bestScore = score
			}
			beta = math.Min(beta, score)
		}

		if beta <= alpha {
			// Store killer move
			if ply < 10 && !ai.isCapture(move, game) {
				ai.killerMoves[ply][1] = ai.killerMoves[ply][0]
				ai.killerMoves[ply][0] = move
			}
			// Update history table
			ai.historyTable[move] += depth * depth
			break // Alpha-beta cutoff
		}
	}

	// Store in transposition table
	flag := 0 // exact
	if bestScore <= originalAlpha {
		flag = 2 // upper bound
	} else if bestScore >= beta {
		flag = 1 // lower bound
	}

	ai.transpositionTable[hash] = TranspositionEntry{
		depth: depth,
		score: bestScore,
		flag:  flag,
	}

	return bestScore
}

// orderMoves orders moves for better alpha-beta pruning
func (ai *AI) orderMoves(moves []Move, game *Game, ply int) []Move {
	type scoredMove struct {
		move  Move
		score int
	}

	scoredMoves := make([]scoredMove, len(moves))

	for i, move := range moves {
		score := 0

		// Prioritize captures (MVV-LVA: Most Valuable Victim - Least Valuable Attacker)
		if ai.isCapture(move, game) {
			attacker := game.Board.GetPiece(move.From)
			victim := game.Board.GetPiece(move.To)
			if attacker != nil && victim != nil {
				score += ai.getPieceValue(victim.Type)*10 - ai.getPieceValue(attacker.Type)
			}
		}

		// Prioritize killer moves
		if ply < 10 {
			if move == ai.killerMoves[ply][0] {
				score += 900
			} else if move == ai.killerMoves[ply][1] {
				score += 800
			}
		}

		// History heuristic
		if historyScore, exists := ai.historyTable[move]; exists {
			score += historyScore / 100
		}

		// Prioritize center moves
		if ai.isCenterMove(move.To) {
			score += 50
		}

		// Prioritize piece development
		piece := game.Board.GetPiece(move.From)
		if piece != nil && (piece.Type == Knight || piece.Type == Bishop) {
			if (piece.Color == White && move.From.Row == 0) ||
				(piece.Color == Black && move.From.Row == 7) {
				score += 30
			}
		}

		scoredMoves[i] = scoredMove{move: move, score: score}
	}

	// Sort moves by score (highest first)
	sort.Slice(scoredMoves, func(i, j int) bool {
		return scoredMoves[i].score > scoredMoves[j].score
	})

	// Extract sorted moves
	result := make([]Move, len(moves))
	for i, sm := range scoredMoves {
		result[i] = sm.move
	}

	return result
}

// evaluatePosition evaluates the current position and returns a score (enhanced)
func (ai *AI) evaluatePosition(game *Game) float64 {
	if game.State == Checkmate {
		if game.CurrentPlayer == ai.color {
			return -1000.0 // AI is checkmated
		}
		return 1000.0 // Opponent is checkmated
	}

	if game.State == Stalemate {
		return 0.0 // Draw
	}

	score := 0.0

	// Enhanced piece values
	pieceValues := map[PieceType]float64{
		Pawn:   1.0,
		Knight: 3.2,
		Bishop: 3.3,
		Rook:   5.0,
		Queen:  9.0,
		King:   0.0,
	}

	// Enhanced position tables
	pawnTable := [8][8]float64{
		{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8},
		{0.3, 0.3, 0.4, 0.5, 0.5, 0.4, 0.3, 0.3},
		{0.1, 0.1, 0.2, 0.4, 0.4, 0.2, 0.1, 0.1},
		{0.05, 0.05, 0.1, 0.3, 0.3, 0.1, 0.05, 0.05},
		{0.05, -0.05, -0.1, 0.0, 0.0, -0.1, -0.05, 0.05},
		{0.05, 0.1, 0.1, -0.2, -0.2, 0.1, 0.1, 0.05},
		{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	knightTable := [8][8]float64{
		{-0.5, -0.4, -0.3, -0.3, -0.3, -0.3, -0.4, -0.5},
		{-0.4, -0.2, 0.0, 0.0, 0.0, 0.0, -0.2, -0.4},
		{-0.3, 0.0, 0.1, 0.15, 0.15, 0.1, 0.0, -0.3},
		{-0.3, 0.05, 0.15, 0.2, 0.2, 0.15, 0.05, -0.3},
		{-0.3, 0.0, 0.15, 0.2, 0.2, 0.15, 0.0, -0.3},
		{-0.3, 0.05, 0.1, 0.15, 0.15, 0.1, 0.05, -0.3},
		{-0.4, -0.2, 0.0, 0.05, 0.05, 0.0, -0.2, -0.4},
		{-0.5, -0.4, -0.3, -0.3, -0.3, -0.3, -0.4, -0.5},
	}

	bishopTable := [8][8]float64{
		{-0.2, -0.1, -0.1, -0.1, -0.1, -0.1, -0.1, -0.2},
		{-0.1, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.1},
		{-0.1, 0.0, 0.05, 0.1, 0.1, 0.05, 0.0, -0.1},
		{-0.1, 0.05, 0.05, 0.1, 0.1, 0.05, 0.05, -0.1},
		{-0.1, 0.0, 0.1, 0.1, 0.1, 0.1, 0.0, -0.1},
		{-0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, -0.1},
		{-0.1, 0.05, 0.0, 0.0, 0.0, 0.0, 0.05, -0.1},
		{-0.2, -0.1, -0.1, -0.1, -0.1, -0.1, -0.1, -0.2},
	}

	rookTable := [8][8]float64{
		{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{0.05, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.05},
		{-0.05, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.05},
		{-0.05, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.05},
		{-0.05, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.05},
		{-0.05, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.05},
		{-0.05, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.05},
		{0.0, 0.0, 0.0, 0.05, 0.05, 0.0, 0.0, 0.0},
	}

	// Material and positional evaluation
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := game.Board.GetPiece(NewPosition(row, col))
			if piece == nil {
				continue
			}

			value := pieceValues[piece.Type]

			// Add positional bonuses
			switch piece.Type {
			case Pawn:
				if piece.Color == White {
					value += pawnTable[7-row][col]
				} else {
					value += pawnTable[row][col]
				}
			case Knight:
				value += knightTable[row][col]
			case Bishop:
				value += bishopTable[row][col]
			case Rook:
				value += rookTable[row][col]
			}

			// Piece development bonus
			if (piece.Type == Knight || piece.Type == Bishop) &&
				((piece.Color == White && row != 0) || (piece.Color == Black && row != 7)) {
				value += 0.15
			}

			// Apply score based on color
			if piece.Color == ai.color {
				score += value
			} else {
				score -= value
			}
		}
	}

	// Enhanced center control
	centerSquares := []Position{
		NewPosition(3, 3), NewPosition(3, 4),
		NewPosition(4, 3), NewPosition(4, 4),
	}
	extendedCenter := []Position{
		NewPosition(2, 2), NewPosition(2, 3), NewPosition(2, 4), NewPosition(2, 5),
		NewPosition(3, 2), NewPosition(3, 5),
		NewPosition(4, 2), NewPosition(4, 5),
		NewPosition(5, 2), NewPosition(5, 3), NewPosition(5, 4), NewPosition(5, 5),
	}

	for _, pos := range centerSquares {
		piece := game.Board.GetPiece(pos)
		if piece != nil {
			if piece.Color == ai.color {
				score += 0.3
			} else {
				score -= 0.3
			}
		}
	}

	for _, pos := range extendedCenter {
		piece := game.Board.GetPiece(pos)
		if piece != nil {
			if piece.Color == ai.color {
				score += 0.1
			} else {
				score -= 0.1
			}
		}
	}

	// King safety evaluation
	aiKingPos := ai.findKing(game, ai.color)

	if aiKingPos != nil {
		// Penalize exposed king
		if aiKingPos.Row >= 2 && aiKingPos.Row <= 5 && aiKingPos.Col >= 2 && aiKingPos.Col <= 5 {
			score -= 0.8
		}
		// Bonus for king safety behind pawns
		if ai.color == White && aiKingPos.Row <= 1 {
			score += 0.2
		} else if ai.color == Black && aiKingPos.Row >= 6 {
			score += 0.2
		}
	}

	// Mobility bonus
	aiMoves := len(ai.getAllPossibleMovesForColor(game, ai.color))
	opponentMoves := len(ai.getAllPossibleMovesForColor(game, ai.getOpponentColor()))
	score += float64(aiMoves-opponentMoves) * 0.05

	return score
}

// Helper functions for AI optimizations

func (ai *AI) hashPosition(game *Game) uint64 {
	// Simple hash function for transposition table
	hash := uint64(0)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := game.Board.GetPiece(NewPosition(row, col))
			if piece != nil {
				pieceValue := uint64(piece.Type) + uint64(piece.Color)*10
				hash ^= pieceValue * uint64(row*8+col+1)
			}
		}
	}
	hash ^= uint64(game.CurrentPlayer) * 1000000
	return hash
}

func (ai *AI) isCapture(move Move, game *Game) bool {
	return game.Board.GetPiece(move.To) != nil
}

func (ai *AI) isCenterMove(pos Position) bool {
	return (pos.Row >= 2 && pos.Row <= 5) && (pos.Col >= 2 && pos.Col <= 5)
}

func (ai *AI) getPieceValue(pieceType PieceType) int {
	values := map[PieceType]int{
		Pawn: 1, Knight: 3, Bishop: 3, Rook: 5, Queen: 9, King: 0,
	}
	return values[pieceType]
}

func (ai *AI) isInCheck(game *Game, color Color) bool {
	return game.isInCheck(color)
}

func (ai *AI) getAllPossibleMovesForColor(game *Game, color Color) []Move {
	var moves []Move
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			from := NewPosition(row, col)
			piece := game.Board.GetPiece(from)
			if piece == nil || piece.Color != color {
				continue
			}
			validMoves := game.Board.GetValidMoves(from)
			moves = append(moves, validMoves...)
		}
	}
	return moves
}

// getAllPossibleMoves returns all possible moves for the current player
func (ai *AI) getAllPossibleMoves(game *Game) []Move {
	return ai.getAllPossibleMovesForColor(game, game.CurrentPlayer)
}

// copyGame creates a deep copy of the game state
func (ai *AI) copyGame(game *Game) *Game {
	newGame := &Game{
		Board:         NewBoard(),
		CurrentPlayer: game.CurrentPlayer,
		State:         game.State,
		MoveHistory:   make([]Move, len(game.MoveHistory)),
	}

	// Copy the board
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			pos := NewPosition(row, col)
			piece := game.Board.GetPiece(pos)
			if piece != nil {
				newPiece := &Piece{
					Type:  piece.Type,
					Color: piece.Color,
				}
				newGame.Board.SetPiece(pos, newPiece)
			}
		}
	}

	// Copy move history
	copy(newGame.MoveHistory, game.MoveHistory)

	return newGame
}

// findKing finds the king of the specified color
func (ai *AI) findKing(game *Game, color Color) *Position {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			pos := NewPosition(row, col)
			piece := game.Board.GetPiece(pos)
			if piece != nil && piece.Type == King && piece.Color == color {
				return &pos
			}
		}
	}
	return nil
}

// getOpponentColor returns the opposite color
func (ai *AI) getOpponentColor() Color {
	if ai.color == White {
		return Black
	}
	return White
}
