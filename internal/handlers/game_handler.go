package handlers

import (
	"concurrent-game/internal/game"
	"concurrent-game/pkg/ui"
)

// GameHandler gestiona la interacción entre la UI y el juego
type GameHandler struct {
	GameService *game.GameService
}

// NewGameHandler inicializa el manejador del juego
func NewGameHandler() *GameHandler {
	return &GameHandler{
		GameService: game.NewGame(),
	}
}

// Run inicia la interfaz de usuario y el juego
func (gh *GameHandler) Run() {
	ui.RunUI(gh.GameService)
}
