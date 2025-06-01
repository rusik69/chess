// Package main provides the entry point for the chess game application.
package main

import "chess-game/ui"

func main() {
	gameInterface := ui.NewInterface()
	gameInterface.Run()
}
