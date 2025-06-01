# Chess Game in Go

[![CI](https://github.com/rusik69/app/actions/workflows/ci.yml/badge.svg)](https://github.com/rusik69/app/actions/workflows/ci.yml)
[![Release](https://github.com/rusik69/app/actions/workflows/release.yml/badge.svg)](https://github.com/rusik69/app/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rusik69/app)](https://goreportcard.com/report/github.com/rusik69/app)
[![codecov](https://codecov.io/gh/rusik69/app/branch/main/graph/badge.svg)](https://codecov.io/gh/rusik69/app)

A complete chess game implementation in Go with a terminal-based interface and AI opponent.

## Features

- ✅ Complete chess rules implementation
- ✅ All piece movements (King, Queen, Rook, Bishop, Knight, Pawn)
- ✅ Check and checkmate detection
- ✅ Stalemate detection
- ✅ Move validation
- ✅ Turn-based gameplay
- ✅ Beautiful terminal interface
- ✅ Algebraic notation for moves
- ✅ Valid moves display for any piece
- ✅ Game status tracking
- ✅ **AI opponent playing as Black**
- ✅ **Minimax algorithm with alpha-beta pruning**
- ✅ **Position evaluation with piece-square tables**

## How to Run

### Using Makefile (Recommended)

The project includes a comprehensive Makefile for easy development:

```bash
# Run the game directly
make run

# Build and run (default target)
make

# Build the binary
make build

# Run tests
make test

# Run tests with coverage
make test-cover

# Quick development cycle (deps, check, test)
make dev

# See all available commands
make help
```

### Manual Commands

1. Make sure you have Go installed (version 1.24 or later)
2. Clone or download this project
3. Navigate to the project directory
4. Run the game:

```bash
go run main.go
```

Or build and run:

```bash
go build -o bin/chess-game
./bin/chess-game
```

## Makefile Commands

The Makefile provides the following commands:

### Building
- `make build` - Build the chess game binary to `bin/` directory
- `make build-linux` - Cross-compile for Linux (amd64)
- `make build-windows` - Cross-compile for Windows (amd64)
- `make build-mac` - Cross-compile for macOS (amd64)

### Running
- `make run` - Build and run the chess game
- `make dev` - Development workflow (deps, check, test)

### Testing
- `make test` - Run all tests
- `make test-coverage` - Run tests with coverage report
- `make test-cover` - Generate HTML coverage report

### Code Quality
- `make fmt` - Format all Go files
- `make vet` - Run go vet
- `make lint` - Run golangci-lint (if installed)
- `make check` - Run fmt, vet, and lint

### Maintenance
- `make clean` - Remove build artifacts and coverage files
- `make deps` - Download and tidy dependencies

## How to Play

You play as **White** against the computer which plays as **Black**. The computer will automatically make its moves after you make yours.

### Basic Commands

- **Make a move**: Enter moves in the format `<from> <to>`
  - Example: `e2 e4` (moves pawn from e2 to e4)
- **See valid moves**: `moves <position>`
  - Example: `moves e2` (shows all valid moves for piece at e2)
- **Get help**: `help`
- **Quit game**: `quit` or `exit`

### Game Flow

1. **You play as White** - Make the first move
2. **Computer plays as Black** - Automatically responds after your move
3. The computer will display "Computer is thinking..." and then show its move
4. Continue alternating until the game ends

### Board Coordinates

The board uses standard algebraic notation:
- **Files** (columns): a-h (left to right)
- **Ranks** (rows): 1-8 (bottom to top)

### Piece Symbols

- **♔/♚**: King (White/Black)
- **♕/♛**: Queen (White/Black)
- **♖/♜**: Rook (White/Black)
- **♗/♝**: Bishop (White/Black)
- **♘/♞**: Knight (White/Black)
- **♙/♟**: Pawn (White/Black)

## AI Features

The computer opponent includes advanced chess AI techniques:

### **Core Algorithm**
- **Minimax Algorithm**: Searches ahead to find the best moves
- **Alpha-Beta Pruning**: Optimizes search performance by eliminating inferior branches
- **Iterative Deepening**: Gradually increases search depth for better time management
- **Transposition Table**: Caches previously evaluated positions to avoid redundant calculations

### **Move Ordering Optimizations**
- **MVV-LVA (Most Valuable Victim - Least Valuable Attacker)**: Prioritizes captures of valuable pieces
- **Killer Move Heuristic**: Remembers moves that caused cutoffs at each depth
- **History Heuristic**: Tracks historically good moves for better ordering
- **Center Control Priority**: Favors moves that control central squares
- **Development Bonus**: Encourages early piece development

### **Position Evaluation**
- **Enhanced Piece Values**: Fine-tuned piece valuations (Knight: 3.2, Bishop: 3.3)
- **Piece-Square Tables**: Positional bonuses for pawns, knights, bishops, and rooks
- **Center Control**: Bonuses for controlling central and extended central squares
- **King Safety**: Evaluates king position and pawn shelter
- **Piece Development**: Encourages moving pieces from starting positions
- **Mobility**: Considers the number of available moves for each side

### **Performance Features**
- **Faster Search**: Optimized move ordering reduces search time significantly
- **Smarter Evaluation**: More sophisticated position assessment
- **Memory Efficiency**: Transposition table prevents redundant calculations
- **Tactical Awareness**: Better at finding captures and threats

The AI plays at depth 3 with iterative deepening, providing strong tactical play while maintaining reasonable response times. The enhanced move ordering and evaluation make it significantly stronger than the basic version.

## Example Gameplay

```
   a  b  c  d  e  f  g  h
  ┌──┬──┬──┬──┬──┬──┬──┬──┐
8 │♜ │♞ │♝ │♛ │♚ │♝ │♞ │♜ │ 8
  ├──┼──┼──┼──┼──┼──┼──┼──┤
7 │♟ │♟ │♟ │♟ │♟ │♟ │♟ │♟ │ 7
  ├──┼──┼──┼──┼──┼──┼──┼──┤
6 │  │  │  │  │  │  │  │  │ 6
  ├──┼──┼──┼──┼──┼──┼──┼──┤
5 │  │  │  │  │  │  │  │  │ 5
  ├──┼──┼──┼──┼──┼──┼──┼──┤
4 │  │  │  │  │♙ │  │  │  │ 4
  ├──┼──┼──┼──┼──┼──┼──┼──┤
3 │  │  │  │  │  │  │  │  │ 3
  ├──┼──┼──┼──┼──┼──┼──┼──┤
2 │♙ │♙ │♙ │♙ │  │♙ │♙ │♙ │ 2
  ├──┼──┼──┼──┼──┼──┼──┼──┤
1 │♖ │♘ │♗ │♕ │♔ │♗ │♘ │♖ │ 1
  └──┴──┴──┴──┴──┴──┴──┴──┘
   a  b  c  d  e  f  g  h

Current Player: Black
Game State: Playing

Computer is thinking...
Computer plays: g8 -> f6
```

## Project Structure

```
chess-game/
├── main.go              # Main entry point
├── go.mod              # Go module file
├── Makefile            # Build automation
├── README.md           # This file
├── .gitignore          # Git ignore file
├── bin/                # Compiled binaries (created by build)
├── chess/              # Chess game logic
│   ├── piece.go        # Piece types and definitions
│   ├── position.go     # Board position handling
│   ├── board.go        # Chess board implementation
│   ├── moves.go        # Move validation logic
│   ├── game.go         # Game state management
│   ├── ai.go           # AI engine with minimax algorithm
│   └── game_test.go    # Unit tests
└── ui/                 # User interface
    ├── interface.go    # Terminal-based interface with AI integration
    └── interface_test.go # UI unit tests
```

## Development

### Quick Start
```bash
# Clone the repository
git clone <repository-url>
cd chess-game

# Run tests and build
make

# Start playing against the computer
make run
```

### Development Workflow
```bash
# Quick development cycle
make dev

# Check code quality
make check

# Run tests with coverage
make test-cover
```

## Game Rules Implemented

### Piece Movement
- **Pawn**: Moves forward one square, two squares from starting position, captures diagonally
- **Rook**: Moves horizontally and vertically any number of squares
- **Knight**: Moves in L-shape (2+1 squares)
- **Bishop**: Moves diagonally any number of squares
- **Queen**: Combines rook and bishop movement
- **King**: Moves one square in any direction

### Special Rules
- **Check**: When a king is under attack
- **Checkmate**: When a king is in check and has no legal moves
- **Stalemate**: When a player has no legal moves but is not in check

## Future Enhancements

Potential features that could be added:
- Castling
- En passant capture
- Pawn promotion
- Draw by repetition
- 50-move rule
- PGN (Portable Game Notation) support
- Adjustable AI difficulty levels
- Opening book for AI
- Endgame tablebase support
- Network multiplayer
- GUI interface

## CI/CD Pipeline

This project uses GitHub Actions for continuous integration and deployment:

### **Continuous Integration**
- **Go 1.24 Testing**: Tests against the latest Go version
- **Cross-platform Builds**: Builds for Linux, macOS, and Windows on AMD64 and ARM64
- **Code Quality Checks**: 
  - `go vet` for static analysis
  - `gofmt` for code formatting
  - `golangci-lint` for comprehensive linting
  - Race condition detection with `-race` flag
- **Security Scanning**: 
  - Gosec for security vulnerabilities
  - Nancy for dependency vulnerabilities
- **Test Coverage**: Generates coverage reports and uploads to Codecov

### **Automated Releases**
- **Tagged Releases**: Automatically creates releases when tags are pushed
- **Cross-platform Binaries**: Builds and packages binaries for all supported platforms
- **Release Notes**: Auto-generates release notes with download links
- **Artifact Storage**: Stores build artifacts for easy access

### **Workflow Triggers**
- **CI Pipeline**: Runs on push to `main`, `master`, `develop` branches and all pull requests
- **Release Pipeline**: Runs on version tags (e.g., `v1.0.0`)

### **Quality Gates**
- All tests must pass before builds are created
- Code must pass linting and security checks
- Coverage reports are generated for all builds

### **Action Versions**
All workflows use the latest stable versions of GitHub Actions:
- `actions/checkout@v4` - Repository checkout
- `actions/setup-go@v4` - Go environment setup
- `actions/cache@v4` - Dependency caching
- `actions/upload-artifact@v4` - Artifact storage
- `golangci/golangci-lint-action@v4` - Code linting
- `codecov/codecov-action@v4` - Coverage reporting
- `softprops/action-gh-release@v2` - Release creation

To create a new release:
```bash
git tag v1.0.0
git push origin v1.0.0
```

## Contributing

Feel free to contribute to this project by:
1. Reporting bugs
2. Suggesting new features
3. Submitting pull requests
4. Improving documentation
5. Enhancing the AI algorithm

## License

This project is open source and available under the MIT License. 