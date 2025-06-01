package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"chess-game/chess"
)

// Interface handles the user interface for the chess game
type Interface struct {
	game   *chess.Game
	reader *bufio.Reader
	ai     *chess.AI
}

// NewInterface creates a new interface
func NewInterface() *Interface {
	return &Interface{
		game:   chess.NewGame(),
		reader: bufio.NewReader(os.Stdin),
		ai:     chess.NewAI(chess.Black, 3), // AI plays as black with depth 3
	}
}

// clearScreen clears the terminal screen
func (ui *Interface) clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// displayWelcome displays the welcome message
func (ui *Interface) displayWelcome() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║          Welcome to Chess!          ║")
	fmt.Println("║                                      ║")
	fmt.Println("║    You play as White vs Computer     ║")
	fmt.Println("║                                      ║")
	fmt.Println("║  Enter moves in algebraic notation   ║")
	fmt.Println("║  Example: e2 e4 (move from e2 to e4) ║")
	fmt.Println("║                                      ║")
	fmt.Println("║  Commands:                           ║")
	fmt.Println("║  - 'quit' or 'exit' to quit          ║")
	fmt.Println("║  - 'help' for help                   ║")
	fmt.Println("║  - 'moves <pos>' to see valid moves  ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Println()
}

// displayHelp displays help information
func (ui *Interface) displayHelp() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                Help                  ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║ Piece Symbols:                       ║")
	fmt.Println("║ ♔/♚ = King    ♕/♛ = Queen           ║")
	fmt.Println("║ ♖/♜ = Rook    ♗/♝ = Bishop          ║")
	fmt.Println("║ ♘/♞ = Knight  ♙/♟ = Pawn            ║")
	fmt.Println("║ (White/Black pieces)                 ║")
	fmt.Println("║                                      ║")
	fmt.Println("║ Move Format:                         ║")
	fmt.Println("║ <from> <to>                          ║")
	fmt.Println("║ Example: e2 e4                       ║")
	fmt.Println("║                                      ║")
	fmt.Println("║ Board Coordinates:                   ║")
	fmt.Println("║ Files: a-h (left to right)          ║")
	fmt.Println("║ Ranks: 1-8 (bottom to top)          ║")
	fmt.Println("║                                      ║")
	fmt.Println("║ You are playing as White             ║")
	fmt.Println("║ Computer is playing as Black         ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Println()
}

// displayBoard displays the current board state
func (ui *Interface) displayBoard() {
	fmt.Println(ui.game.Board.String())
}

// displayGameStatus displays the current game status
func (ui *Interface) displayGameStatus() {
	fmt.Println(ui.game.GetGameStatus())
}

// getInput gets input from the user
func (ui *Interface) getInput() string {
	fmt.Printf("%s> ", ui.game.CurrentPlayer)
	input, _ := ui.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// showValidMoves displays valid moves for a piece at the given position
func (ui *Interface) showValidMoves(position string) {
	pos, err := chess.FromAlgebraic(position)
	if err != nil {
		fmt.Printf("Invalid position: %s\n", position)
		return
	}

	piece := ui.game.Board.GetPiece(pos)
	if piece == nil {
		fmt.Printf("No piece at position %s\n", position)
		return
	}

	if piece.Color != ui.game.CurrentPlayer {
		fmt.Printf("That's not your piece!\n")
		return
	}

	moves := ui.game.Board.GetValidMoves(pos)
	if len(moves) == 0 {
		fmt.Printf("No valid moves for piece at %s\n", position)
		return
	}

	fmt.Printf("Valid moves for %s %s at %s:\n", piece.Color, piece.Type, position)
	for i, move := range moves {
		fmt.Printf("%d. %s", i+1, move.To)
		if (i+1)%8 == 0 {
			fmt.Println()
		} else {
			fmt.Print("  ")
		}
	}
	if len(moves)%8 != 0 {
		fmt.Println()
	}
	fmt.Println()
}

// processMove processes a move input
func (ui *Interface) processMove(input string) bool {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Invalid move format. Use: <from> <to> (e.g., e2 e4)")
		return false
	}

	from := parts[0]
	to := parts[1]

	err := ui.game.MakeMove(from, to)
	if err != nil {
		fmt.Printf("Invalid move: %v\n", err)
		return false
	}

	return true
}

// makeAIMove makes a move for the AI player
func (ui *Interface) makeAIMove() bool {
	if ui.game.CurrentPlayer != chess.Black {
		return false
	}

	fmt.Println("Computer is thinking...")

	// Add a small delay to make it feel more natural
	time.Sleep(1 * time.Second)

	from, to, ok := ui.ai.GetBestMove(ui.game)
	if !ok {
		fmt.Println("Computer has no valid moves!")
		return false
	}

	err := ui.game.MakeMove(from.String(), to.String())
	if err != nil {
		fmt.Printf("Computer move error: %v\n", err)
		return false
	}

	fmt.Printf("Computer plays: %s -> %s\n", from.String(), to.String())
	return true
}

// Run starts the chess game interface
func (ui *Interface) Run() {
	ui.clearScreen()
	ui.displayWelcome()

	for !ui.game.IsGameOver() {
		ui.displayBoard()
		ui.displayGameStatus()

		// Check if it's the AI's turn
		if ui.game.CurrentPlayer == chess.Black {
			if !ui.makeAIMove() {
				break
			}
			ui.clearScreen()
			continue
		}

		// Human player's turn
		input := ui.getInput()

		// Handle commands
		switch {
		case input == "quit" || input == "exit":
			fmt.Println("Thanks for playing!")
			return
		case input == "help":
			ui.displayHelp()
			continue
		case strings.HasPrefix(input, "moves "):
			position := strings.TrimPrefix(input, "moves ")
			ui.showValidMoves(position)
			continue
		case input == "":
			continue
		}

		// Process human move
		if ui.processMove(input) {
			ui.clearScreen()
		}
	}

	// Game over
	ui.displayBoard()
	ui.displayGameStatus()
	fmt.Println("Game Over!")
}
