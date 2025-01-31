package main

import "testing"

func TestPlayerSelectionPure(t *testing.T) {
	selectPlayerOne, err := PlayerSelectPure(0)
	if selectPlayerOne != PlayerOne || err != nil {
		t.Errorf("expected %v but got %v\n", PlayerOne, selectPlayerOne)
	}
	selectPlayerTwo, err := PlayerSelectPure(1)
	if selectPlayerTwo != PlayerTwo || err != nil {
		t.Errorf("expected %v but go %v\n", PlayerOne, selectPlayerTwo)
	}
	_, err = PlayerSelectPure(2)
	if err != nil {
		t.Error("expected error but received nil", err)
	}
}
