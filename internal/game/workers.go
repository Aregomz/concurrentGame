package game

import (
	"sync"
	"time"
	"concurrent-game/internal/models"
)

// SimulateRace es un worker que simula el progreso de los jugadores.
func SimulateRace(player *models.Player, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for !player.Finished() {
		mu.Lock()
		player.Steps += 1
		mu.Unlock()
		time.Sleep(time.Millisecond * 100)  // Simula el tiempo de carrera
	}
}
