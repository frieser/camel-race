package player

import (
	"github.com/frieser/camelrace/camel"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/terrain"
	"testing"
)

func mockCpu() Cpu {
	cpu := Cpu{
		Name:      "Cpu Player",
		Transport: camel.NewDomesticCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_GRASS}}, event.EventDispatcher{}),
		Distance:  0}

	return cpu
}

func TestCpuRoll(t *testing.T) {
	cpu := mockCpu()

	roll := cpu.Roll()

	if roll < 0 || roll > 3 {
		t.Error("Cpu Roll is not between the race ranges")
	}
}

func TestCpuPlay(t *testing.T) {
	cpu := mockCpu()
	cpu.Play()

	if cpu.Distance < 0 || cpu.Distance > 3 {
		t.Error("Cpu Roll is not between the race ranges")
	}
}
