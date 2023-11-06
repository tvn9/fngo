// Let's refactor the code into something testable
package main

import (
	"fmt"
	"math/rand"
)

type Player string

const (
	PlayerOne Player = "Remi"
	PlayerTwo Player = "Yvonne"
)

// refacture into something testable
func PlayerSelectPure(i int) (Player, error) {
	switch i {
	case 0:
		return PlayerOne, nil
	case 1:
		return PlayerTwo, nil
	}
	return Player(""),
		fmt.Errorf("no player matching input: %v", i)
}

// randomized the input value, result in random result
func selectStartingPlayer() Player {
	randomized := rand.Intn(2)
	switch randomized {
	case 0:
		return PlayerOne
	case 1:
		return PlayerTwo
	}
	panic("No further player available")
}

func main() {
	p := selectStartingPlayer()
	fmt.Println(p)
	p1, err := PlayerSelectPure(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)

}
