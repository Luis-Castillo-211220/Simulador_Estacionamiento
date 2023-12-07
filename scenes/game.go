package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"

	"Simulator/controllers"
	"Simulator/models"
)

type GameScene struct {
	window  fyne.Window
	content *fyne.Container
}

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
	scene.Render()
	return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))

	s.content = container.NewWithoutLayout(backgroundImage)
	s.window.SetContent(s.content)
	s.StartGame()
}

func (s *GameScene) StartGame() {
	var GenerateCarros = controllers.InicializarCarros
	e := models.CreateEstacionamiento(100)
	go GenerateCarros(100, e)
	go s.PintarCarros(e)
}
