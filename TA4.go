package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	TamanioTablero = 58
	numJugadores   = 4
	fichasMax      = 4
	PosWin         = 57
)

var (
	tablero [TamanioTablero]int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

type Jugador struct {
	ID             int
	fichas         [fichasMax]int
	ganador        bool
	siguienteFicha int
	turno          chan bool
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < TamanioTablero; i++ {
		tablero[i] = -1
	}

	jugadores := make([]Jugador, numJugadores)
	for i := 0; i < numJugadores; i++ {
		jugadores[i] = Jugador{ID: i, turno: make(chan bool)}
	}

	wg.Add(numJugadores)

	for i := 0; i < numJugadores; i++ {
		go jugarLudo(&jugadores[i], jugadores)
	}

	jugadores[0].turno <- true

	wg.Wait()
}

func jugarLudo(jugador *Jugador, jugadores []Jugador) {
	defer wg.Done()

	for !jugador.ganador {
		<-jugador.turno

		dados := tirarTresDados()

		mu.Lock()
		if jugador.fichas[jugador.siguienteFicha] >= PosWin {
			jugador.siguienteFicha++
			if jugador.siguienteFicha == fichasMax {
				jugador.ganador = true
				fmt.Printf("JUGADOR %d HA GANADO!!!!!!!\n", jugador.ID+1)
				os.Exit(0)
			}
			mu.Unlock()
			continue
		}

		fichaPos := jugador.fichas[jugador.siguienteFicha]

		var nuevaPos int
		if fichaPos+dados[2] < 0 {
			nuevaPos = 0
		} else {
			nuevaPos = fichaPos + dados[2]
		}

		if nuevaPos > PosWin {
			nuevaPos = PosWin
		}

		jugador.fichas[jugador.siguienteFicha] = nuevaPos

		fmt.Print("*******************************************************************************************************\n")
		fmt.Printf("TURNO DEL JUGADOR %d\n", jugador.ID+1)

		if dados[3] == 0 {
			fmt.Printf("<<<<Jugador %d, ha sacado un + y en este turno AVANZARA>>>>.\n", jugador.ID+1)
		} else {
			fmt.Printf("<<<<Jugador %d, ha sacado un - y en este turno RETROCEDERA>>>>.\n", jugador.ID+1)
		}

		fmt.Printf("Jugador %d, valor del primer dado: %d, valor del segundo dado: %d, CANTIDAD DE CASILLAS A MOVERSE:%d.\n", jugador.ID+1, dados[0], dados[1], dados[2])
		fmt.Printf("Jugador %d, peon número %d, posicion %d.\n", jugador.ID+1, jugador.siguienteFicha+1, nuevaPos)
		fmt.Print("*******************************************************************************************************\n\n")

		if casillaEspecial(nuevaPos) {
			fmt.Printf("Jugador %d pierde un turno, cayo en una casilla multiplo de 10\n\n", jugador.ID+1)
		}

		if nuevaPos >= PosWin {
			jugador.siguienteFicha++
			if jugador.siguienteFicha == fichasMax {
				jugador.ganador = true
				fmt.Printf("JUGADOR %d HA GANADO!!!!!!!!\n", jugador.ID+1)
				os.Exit(0)
			}
		}

		mu.Unlock()

		jugadores[(jugador.ID+1)%numJugadores].turno <- true
	}
}

func tirarTresDados() [4]int {
	dado1 := rand.Intn(6) + 1
	dado2 := rand.Intn(6) + 1

	operacion := rand.Intn(2)
	pasos := 0

	if operacion == 0 {
		pasos = dado1 + dado2
	} else {
		pasos = (dado1 + dado2) * -1
	}

	return [4]int{dado1, dado2, pasos, operacion}
}

func casillaEspecial(fichaPos int) bool {
	return fichaPos%10 == 0 && fichaPos != 0
}
