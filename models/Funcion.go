package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
)

func (c *Carro) RunCarro() {
	for {
		// Verificar si hay espacio disponible antes de intentar estacionar
		if c.Estacionamiento.EspaciosDisponibles > 0 {
			c.estacionar()
			time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
			c.abandonar()
		} else {
			// Esperar si no hay espacio disponible
			c.esperarEspacioDisponible()
		}
	}
}

func (c *Carro) estacionar() {
	c.Estacionamiento.Mutex.Lock()
	defer c.Estacionamiento.Mutex.Unlock()
	for len(c.Estacionamiento.EspaciosEstacionamiento) == 0 {
		c.Estacionamiento.Cond.Wait()
	}

	slot := <-c.Estacionamiento.EspaciosEstacionamiento

	if slot {
		x := float32(rand.Intn(650-150+1) + 150)
		y := float32(rand.Intn(300-50+1) + 50)
		c.skin.Move(fyne.NewPos(x, y))
		fmt.Println("Carro", c.I, "Entra al estacionamiento")

	}

	c.Estacionamiento.EspaciosOcupados++
	c.Estacionamiento.EspaciosDisponibles--

	fmt.Println("Espacios Ocupados: ", c.Estacionamiento.EspaciosOcupados)
	fmt.Println("Espacios Disponibles: ", c.Estacionamiento.EspaciosDisponibles)

	

	// if c.Estacionamiento.EspaciosDisponibles == 0 {
	// 	// Enviar señal para detener la generación de carros
	// 	c.Estacionamiento.DetenerGeneracion <- struct{}{}
	// }

	// c.Estacionamiento.M.Unlock()
}

func (c *Carro) abandonar() {
	c.Estacionamiento.Mutex.Lock()
	defer c.Estacionamiento.Mutex.Unlock()

	c.skin.Move(fyne.NewPos(0, 0))
	fmt.Println("Carro", c.I, "Abandona el estacionamiento")
	c.Estacionamiento.CarroAbandona <- struct{}{}
	c.Estacionamiento.EspaciosEstacionamiento <- true // Enviar true para indicar que el slot está disponible

	// Notificar el cambio en la condición después de haber liberado el candado
	c.Estacionamiento.Cond.Signal()

	c.Estacionamiento.EspaciosOcupados--
	c.Estacionamiento.EspaciosDisponibles++

	fmt.Println("Espacios Ocupados: ", c.Estacionamiento.EspaciosOcupados)
	fmt.Println("Espacios Disponibles: ", c.Estacionamiento.EspaciosDisponibles)

	// c.Estacionamiento.EspacioLiberado <- struct{}{}
	
	// c.Estacionamiento.ReanudarGeneracion <- struct{}{}

	// c.Estacionamiento.M.Unlock()
}

func (c *Carro) esperarEspacioDisponible() {
	fmt.Println("Carro", c.I, "esperando por un espacio disponible.")
	time.Sleep(time.Duration(4* time.Second))
}
