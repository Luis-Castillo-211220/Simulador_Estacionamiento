package models

import (
	"fyne.io/fyne/v2/canvas"
)

type Carro struct {
	Estacionamiento *Estacionamiento
	I  int
	skin *canvas.Image
}

func CreateCarro(e *Estacionamiento, s *canvas.Image) *Carro {
	return &Carro{
		Estacionamiento: e,
		skin: s,
	}
}
