package main

import "math/rand"

type Player string

const (
	PlayerOne Player = "Remi"
	PlayerTwo Player = "Yvonne"
)

func selectStartingPlayer() Player {
	randomized := rand.Intn(2)
	switch randomized {
	case 0:
		return PlayerOne
	case 1:
		return PlayerTwo
	}
	panic("No futher player available")
}
