package controllers

import (
	"Simulator/models"
	// "fmt"
	// "math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

var (
	entranceSemaphore = &sync.Mutex{}
	generarCarros     = true // Variable para controlar la generación de carros
)


func GenerarEspaciosEstacionamiento(filas, columnas, espacioX, espacioY int) []fyne.Position {
	espaciosEstacionamiento := make([]fyne.Position, 0, filas*columnas)

	for fila := 0; fila < filas; fila++ {
		for columna := 0; columna < columnas; columna++ {
			posX := float32(100 + espacioX*columna)
			posY := float32(100 + espacioY*fila)
			espaciosEstacionamiento = append(espaciosEstacionamiento, fyne.NewPos(posX, posY))
		}
	}

	return espaciosEstacionamiento
}


// func generarCarro(estacionamiento *models.Estacionamiento, espaciosEstacionamiento []fyne.Position, generatedCars int) {
//     for {
//         select {
//         case <-estacionamiento.DetenerGeneracion:
//             <-estacionamiento.ReanudarGeneracion
//         case estacionamiento.SlotsEstacionamiento <- true:
//             carroImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/carro1.png"))
//             carroImage.Resize(fyne.NewSize(100, 100))
//             carroImage.Move(espaciosEstacionamiento[generatedCars%len(espaciosEstacionamiento)])
// 			fmt.Println("crea la img")
//             nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
//             nuevoCarro.I = generatedCars + 1

//             estacionamiento.PintarCarro <- carroImage
//             go nuevoCarro.RunCarro()
// 			fmt.Println("crea la go routne del carro")

//             time.Sleep(4 * time.Second)
//         default:
//             <-estacionamiento.VehiculosBloqueados
//         }
//     }
// }


func generarCarro(estacionamiento *models.Estacionamiento, espaciosEstacionamiento []fyne.Position, generatedCars int) {
    select {
    case estacionamiento.EspaciosEstacionamiento <- true:
        carroImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/carro1.png"))
        carroImage.Resize(fyne.NewSize(100, 100))
        carroImage.Move(espaciosEstacionamiento[generatedCars%len(espaciosEstacionamiento)])

        nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
        nuevoCarro.I = generatedCars + 1

        estacionamiento.ImagenCarro <- carroImage
        go nuevoCarro.RunCarro()

        // tiempoEsperar := rand.Intn(5000-1000+1) + 1000
        time.Sleep(4* time.Second)
    default:
        <-estacionamiento.CarrrosBloq
    }
}

func InicializarCarros(n int, estacionamiento *models.Estacionamiento) {
	generatedCars := 0
	espaciosEstacionamiento := GenerarEspaciosEstacionamiento(2, 10, 100, 100)

	for i := 0; i < n; i++ {
		generarCarro(estacionamiento, espaciosEstacionamiento, generatedCars)
		generatedCars++
	}
}
