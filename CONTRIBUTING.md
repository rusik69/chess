# Contributing to Chess Game

Thank you for your interest in contributing to the Chess Game project! This document provides guidelines and information for contributors.

## Development Setup

### Prerequisites
- Go 1.19 or later
- Git
- Make (optional, but recommended)

### Getting Started
1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/app.git
   cd app
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run tests to ensure everything works:
   ```bash
   make test
   ```

## Development Workflow

### Making Changes
1. Create a new branch for your feature/fix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Make your changes
3. Run the development checks:
   ```bash
   make dev  # Runs deps, check, and test
   ```
4. Commit your changes with a descriptive message
5. Push to your fork and create a pull request

### Code Quality Standards

Our CI pipeline enforces several quality standards:

#### **Formatting and Style**
- Code must be formatted with `gofmt`
- Run `make fmt` to format your code
- Follow Go naming conventions

#### **Testing**
- All new code must include tests
- Tests must pass with race detection: `go test -race ./...`
- Aim for high test coverage (current: 97%+)
- Test files should end with `_test.go`

#### **Linting**
- Code must pass `golangci-lint` checks
- Run `make lint` to check locally
- See `.golangci.yml` for configuration

#### **Static Analysis**
- Code must pass `go vet` checks
- Run `make vet` to check locally

### Project Structure

```
chess-game/
├── .github/workflows/    # GitHub Actions CI/CD
├── chess/               # Core chess logic
├── ui/                  # User interface
├── bin/                 # Build artifacts (gitignored)
├── main.go             # Application entry point
├── Makefile            # Build automation
├── .golangci.yml       # Linting configuration
└── README.md           # Project documentation
```

## CI/CD Pipeline

### Continuous Integration
Every push and pull request triggers:

1. **Multi-version Testing**: Go 1.19, 1.20, 1.21
2. **Code Quality Checks**:
   - `go vet` static analysis
   - `gofmt` formatting check
   - `golangci-lint` comprehensive linting
   - Race condition detection
3. **Security Scanning**:
   - Gosec for security vulnerabilities
   - Nancy for dependency vulnerabilities
4. **Cross-platform Builds**: Linux, macOS, Windows (AMD64/ARM64)
5. **Test Coverage**: Reports uploaded to Codecov

### Release Process
1. Create and push a version tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
2. GitHub Actions automatically:
   - Builds cross-platform binaries
   - Creates a GitHub release
   - Uploads packaged artifacts
   - Generates release notes

## Types of Contributions

### Bug Reports
- Use the GitHub issue template
- Include steps to reproduce
- Provide system information (OS, Go version)
- Include relevant logs or error messages

### Feature Requests
- Describe the feature and its use case
- Explain why it would be valuable
- Consider implementation complexity

### Code Contributions
- Bug fixes
- New features
- Performance improvements
- Documentation improvements
- Test coverage improvements

### Chess Engine Improvements
- AI algorithm enhancements
- Position evaluation improvements
- Move ordering optimizations
- Opening book additions
- Endgame tablebase support

## Testing Guidelines

### Unit Tests
- Test all public functions
- Test edge cases and error conditions
- Use table-driven tests where appropriate
- Mock external dependencies

### Integration Tests
- Test complete game scenarios
- Test AI behavior
- Test user interface interactions

### Performance Tests
- Benchmark critical algorithms
- Test AI search performance
- Memory usage optimization

## Code Review Process

1. All changes require a pull request
2. CI checks must pass
3. Code review by maintainers
4. Approval required before merge
5. Squash and merge preferred

## Documentation

- Update README.md for user-facing changes
- Add inline comments for complex logic
- Update this CONTRIBUTING.md for process changes
- Include examples in documentation

## Getting Help

- Open an issue for questions
- Check existing issues and documentation
- Join discussions in pull requests
- Follow Go community best practices

## Recognition

Contributors will be recognized in:
- GitHub contributors list
- Release notes for significant contributions
- README acknowledgments

Thank you for contributing to make this chess game better! 🎉 