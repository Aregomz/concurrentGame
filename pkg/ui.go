package ui

import (
	
    "concurrent-game/internal/game"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "sync"
    "time"
)

// RunUI inicia la interfaz gráfica
func RunUI(g *game.GameService) {
	a := app.New()
	w := a.NewWindow("Carrera de Tortugas")

	// Elementos de la interfaz
	status := widget.NewLabel("Carrera en progreso...")
	playerLabels := []*widget.Label{
		widget.NewLabel("Tortuga 1: 0 pasos"),
		widget.NewLabel("Tortuga 2: 0 pasos"),
		widget.NewLabel("Tortuga 3: 0 pasos"),
	}

	// Contenedor principal
	w.SetContent(container.NewVBox(
		playerLabels[0],
		playerLabels[1],
		playerLabels[2],
		status,
	))

	// Iniciar la carrera en goroutines
	wg := &sync.WaitGroup{}
	go func() {
		g.Start()
		for !g.Game.Finished {
			time.Sleep(100 * time.Millisecond)
			for i, player := range g.Game.Players {
				playerLabels[i].SetText(player.Name + ": " + string(player.Steps) + " pasos")
			}
		}
		status.SetText("¡Ganador: " + g.Game.Winner + "!")
	}()

	// Mostrar la ventana
	w.ShowAndRun()
}
