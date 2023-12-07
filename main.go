package main

import (
	"Simulator/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("EstacionamientoGO")

	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(800, 600))
	scenes.NewGameScene(w)
	w.ShowAndRun()
}
