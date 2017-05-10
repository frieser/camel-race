package camel

import (
	"fmt"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/terrain"
)

type BactrianCamel camel

func NewBactrianCamel(lane lane.Lane, eventDispatcher event.EventDispatcher) *BactrianCamel {
	camel := &BactrianCamel{
		PreferredTerrain: terrain.MATERIAL_MUD,
		Lane:             lane}

	camel.Attributes = append(camel.Attributes, Attribute{Name: "Turbo Boost", Modifier: camel.turboBoostModifier})
	camel.Attributes = append(camel.Attributes, Attribute{Name: "Weapon", Modifier: camel.weaponModifier})
	camel.Attributes = append(camel.Attributes, Attribute{Name: "Strength", Modifier: camel.strengthModifier})

	camel.eventDispatcher = eventDispatcher

	return camel
}

func (camel BactrianCamel) Move(n float32) float32 {
	for _, attr := range camel.Attributes {
		n = (attr.Modifier(n)).(float32)
	}

	return n
}

func (camel BactrianCamel) turboBoostModifier(n float32) interface{} {
	conditionToActivateAttribute := true

	if conditionToActivateAttribute {
		camel.eventDispatcher.Dispatch(event.TurboBoostEvent, camel)

		if camel.Lane.Terrain.Material == camel.PreferredTerrain {
			return n * 1.3
		}

		return n * 1.15
	}

	return n
}

func (camel BactrianCamel) weaponModifier(n float32) interface{} {
	conditionToActivateAttribute := true

	if conditionToActivateAttribute {
		camel.eventDispatcher.Dispatch(event.WeaponEvent, camel)
	}

	return n
}

func (camel BactrianCamel) strengthModifier(n float32) interface{} {
	conditionWhenHitByOtherCamelsWeapons := true

	if conditionWhenHitByOtherCamelsWeapons {
		camel.eventDispatcher.Dispatch(event.StrengthEvent, camel)
		return n * (1 + 0.2)
	}

	return n
}

func (camel BactrianCamel) String() string {
	return fmt.Sprintf("Bactrian Camel")
}
