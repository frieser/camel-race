package player

import (
	"github.com/frieser/camelrace/camel"
	"testing"
)

func TestPlayer(t *testing.T) {
	var distance float32 = 10
	name := "Test player"
	transport := camel.DomesticCamel{}

	player := Player{Name: name, Distance: distance, Transport: transport}

	if player.Name != name {
		t.Error("Bad Player Instantiation")
	}

	if player.Distance != distance {
		t.Error("Bad Player Instantiation")
	}
}
