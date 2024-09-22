package main

import "concurrent-game/internal/handlers"

func main() {
	gameHandler := handlers.NewGameHandler()
	gameHandler.Run()
}
