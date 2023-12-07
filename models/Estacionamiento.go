package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Estacionamiento struct {
	EspaciosEstacionamiento chan bool
	ImagenCarro          chan *canvas.Image
	CarrrosBloq chan struct{}
	CarroAbandona        chan struct{}
	SlotsDisponibles     chan struct{}
	EspaciosOcupados     int
	EspaciosDisponibles  int
	Mutex                    sync.Mutex
	Cond                 *sync.Cond 
	// EspacioLiberado 	 chan struct{}
	DetenerGeneracion 	 chan struct{}
	ReanudarGeneracion 	 chan struct{}
}

func CreateEstacionamiento(nS int) *Estacionamiento {
	return &Estacionamiento{
		EspaciosEstacionamiento: make(chan bool, 100),
		ImagenCarro:          make(chan *canvas.Image, 100),
		CarrrosBloq:  make(chan struct{}),
		CarroAbandona:        make(chan struct{}),
		Cond:                 sync.NewCond(&sync.Mutex{}), 
		SlotsDisponibles:     make(chan struct{}, 20),
		EspaciosOcupados:     (0),
		EspaciosDisponibles:  (20),
		// EspacioLiberado: make(chan struct{}),
		DetenerGeneracion: make(chan struct{}),
		ReanudarGeneracion: make(chan struct{}),
		Mutex: sync.Mutex{},
	}
}
