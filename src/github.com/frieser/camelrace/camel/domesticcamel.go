package camel

import (
	"fmt"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/terrain"
)

type DomesticCamel camel

func NewDomesticCamel(lane lane.Lane, eventDispatcher event.EventDispatcher) *DomesticCamel {
	camel := &DomesticCamel{
		PreferredTerrain: terrain.MATERIAL_GRASS,
		Lane:             lane}

	camel.Attributes = append(camel.Attributes, Attribute{Name: "Turbo Boost", Modifier: camel.turboBoostModifier})
	camel.Attributes = append(camel.Attributes, Attribute{Name: "Weapon", Modifier: camel.weaponModifier})
	camel.Attributes = append(camel.Attributes, Attribute{Name: "Strength", Modifier: camel.strengthModifier})

	camel.eventDispatcher = eventDispatcher

	return camel
}

func (camel DomesticCamel) Move(n float32) float32 {
	for _, attr := range camel.Attributes {
		n = (attr.Modifier(n)).(float32)
	}

	return n
}

func (camel DomesticCamel) turboBoostModifier(n float32) interface{} {
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

func (camel DomesticCamel) weaponModifier(n float32) interface{} {
	conditionToActivateAttribute := true

	if conditionToActivateAttribute {
		camel.eventDispatcher.Dispatch(event.WeaponEvent, camel)
	}

	return n
}

func (camel DomesticCamel) strengthModifier(n float32) interface{} {
	conditionWhenHitByOtherCamelsWeapons := true

	if conditionWhenHitByOtherCamelsWeapons {
		camel.eventDispatcher.Dispatch(event.StrengthEvent, camel)

		return n * (1 + 0.1)
	}

	return n
}

func (camel DomesticCamel) String() string {
	return fmt.Sprintf("Domestic Camel")
}
