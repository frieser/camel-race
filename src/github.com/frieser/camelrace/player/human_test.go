package player

import (
	"github.com/frieser/camelrace/camel"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/terrain"
	"testing"
)

func mockHuman() Human {
	human := Human{
		Name:      "Human Player",
		Transport: camel.NewDomesticCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_GRASS}}, event.EventDispatcher{}),
		Distance:  0}

	return human
}

func TestHumanRoll(t *testing.T) {
	human := mockHuman()

	if human.Roll() != 0 {
		t.Error("Human Roll not get the default value when there is not human interaction")
	}
}

func TestHumanPlay(t *testing.T) {
	human := mockHuman()
	human.Play()

	if human.Distance != 0 {
		t.Error("Human Play not get the default value when there is not human interaction")
	}
}
