package models

// Player representa cada tortuga en la carrera
type Player struct {
	Name  string
	Steps int
}

// Game representa el estado general del juego
type Game struct {
	Players  []Player
	Finished bool
	Winner   string
}
