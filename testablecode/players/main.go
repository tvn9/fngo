// Writing testable code with pure function
package main

import (
	"fmt"
	"log"
	"math/rand"
)

// Player of type string
type Player string

const (
	// PlayerOne of type Player
	PlayerOne Player = "Xuan"
	// PlayerTwo of type Player
	PlayerTwo Player = "Thanh"
)

// Function is not easy to test, and needed the refactor
func selectStartingPlayer() Player {
	randomized := rand.Intn(2)
	switch randomized {
	case 0:
		return PlayerOne
	case 1:
		return PlayerTwo
	}
	log.Fatal("No further player available")
	return Player("")
}

// PlayerSelectPure demonstrates the pure funtion concept
func PlayerSelectPure(i int) (Player, error) {
	switch i {
	case 0:
		return PlayerOne, nil
	case 1:
		return PlayerTwo, nil
	}
	return Player(""), fmt.Errorf("no player matching input: %v", i)
}

func main() {
	random := rand.Intn(2)
	player, err := PlayerSelectPure(random)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(player)
}
