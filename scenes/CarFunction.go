	package scenes

	import (
		"sync"
		"Simulator/models"
	)

	func (s *GameScene) PintarCarros(e *models.Estacionamiento) {
		var mu sync.Mutex

		for {
			imagen := <-e.ImagenCarro

			mu.Lock()
			s.content.Add(imagen)
			mu.Unlock()

			<-e.CarroAbandona

			s.window.Canvas().Refresh(s.content)
		}
	}
