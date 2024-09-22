package game

import (
	"math/rand"
	"sync"
	"time"
	"concurrent-game/internal/models"
)

// GameService maneja la lógica del juego
type GameService struct {
	Game  *models.Game
	mu    sync.Mutex
	wg    sync.WaitGroup
}

// NewGame crea un nuevo servicio del juego
func NewGame() *GameService {
	players := []models.Player{
		{Name: "Tortuga 1", Steps: 0},
		{Name: "Tortuga 2", Steps: 0},
		{Name: "Tortuga 3", Steps: 0},
	}
	return &GameService{
		Game: &models.Game{
			Players:  players,
			Finished: false,
			Winner:   "",
		},
	}
}

// Start inicia el juego y lanza los workers para cada jugador
func (g *GameService) Start() {
	workerChan := make(chan *models.Player) // Canal para los workers
	g.wg.Add(len(g.Game.Players))           // Añadimos goroutines al waitgroup

	// Lanzamos los workers
	for i := range g.Game.Players {
		go g.worker(&g.Game.Players[i], workerChan)
	}

	// Proceso que gestiona las actualizaciones de los workers
	go g.processWorkers(workerChan)

	g.wg.Wait() // Esperamos a que terminen todas las goroutines
	close(workerChan)
}

// worker actualiza los pasos de cada jugador y envía los datos al canal
func (g *GameService) worker(player *models.Player, workerChan chan<- *models.Player) {
	defer g.wg.Done()

	for player.Steps < 100 && !g.Game.Finished {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		g.mu.Lock()
		player.Steps += rand.Intn(10)
		g.mu.Unlock()

		workerChan <- player
	}
}

// processWorkers recibe las actualizaciones y determina si hay un ganador
func (g *GameService) processWorkers(workerChan <-chan *models.Player) {
	for player := range workerChan {
		g.mu.Lock()
		if player.Steps >= 100 && !g.Game.Finished {
			g.Game.Finished = true
			g.Game.Winner = player.Name
		}
		g.mu.Unlock()
	}
}
