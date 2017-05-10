package camel

import (
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/terrain"
	"testing"
)

func TestBactrianCamelMoveByNotPreferredTerrainMaterial(t *testing.T) {
	camel := NewBactrianCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_GRASS}}, event.EventDispatcher{})

	if camel.Move(3) != 4.14 {
		t.Error("Camel not move as expected")
	}
}

func TestBactrianCamelMoveByPreferredTerrainMaterial(t *testing.T) {
	camel := NewBactrianCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_MUD}}, event.EventDispatcher{})

	if camel.Move(3) != 4.68 {
		t.Error("Camel not move as expected")
	}
}

func TestDomesticCamelMoveByNotPreferredTerrainMaterial(t *testing.T) {
	camel := NewDomesticCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_MUD}}, event.EventDispatcher{})

	if camel.Move(3) != 3.7949998 {
		t.Error("Camel not move as expected")
	}
}

func TestDomesticCamelMoveByPreferredTerrainMaterial(t *testing.T) {
	camel := NewBactrianCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_GRASS}}, event.EventDispatcher{})

	if camel.Move(3) != 4.14 {
		t.Error("Camel not move as expected")
	}
}

func TestDromedaryMoveByNotPreferredTerrainMaterial(t *testing.T) {
	camel := NewDromedary(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_MUD}}, event.EventDispatcher{})

	if camel.Move(3) != 0.45000002 {
		t.Error("Camel not move as expected")
	}
}

func TestDromedaryMoveByPreferredTerrainMaterial(t *testing.T) {
	camel := NewDromedary(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_SAND}}, event.EventDispatcher{})

	if camel.Move(3) != 0.90000004 {
		t.Error("Camel not move as expected")
	}
}

func TestBactrianCamelMoveByPreferredTerrainMaterialWithAddedAttribute(t *testing.T) {
	camel := NewBactrianCamel(lane.Lane{Terrain: terrain.Terrain{Material: terrain.MATERIAL_MUD}}, event.EventDispatcher{})

	camel.Attributes = append(camel.Attributes, Attribute{Name: "Test", Modifier: func(n float32) interface{} {
		return n * 2
	}})

	if camel.Move(3) != 9.36 {
		t.Error("Camel not move as expected")
	}
}
