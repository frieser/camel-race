package camel

import (
	"fmt"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/terrain"
)

type Dromedary camel

func NewDromedary(lane lane.Lane, eventDispatcher event.EventDispatcher) *Dromedary {
	camel := &Dromedary{
		PreferredTerrain: terrain.MATERIAL_SAND,
		Lane:             lane}

	camel.Attributes = append(camel.Attributes, Attribute{Name: "Turbo Boost", Modifier: camel.turboBoostModifier})
	camel.Attributes = append(camel.Attributes, Attribute{Name: "Weapon", Modifier: camel.weaponModifier})
	camel.Attributes = append(camel.Attributes, Attribute{Name: "Strength", Modifier: camel.strengthModifier})

	camel.eventDispatcher = eventDispatcher

	return camel
}

func (camel Dromedary) Move(n float32) float32 {
	for _, attr := range camel.Attributes {
		n = (attr.Modifier(n)).(float32)
	}

	return n
}

func (camel Dromedary) turboBoostModifier(n float32) interface{} {
	conditionToActivateAttribute := true

	if conditionToActivateAttribute {
		camel.eventDispatcher.Dispatch(event.TurboBoostEvent, camel)

		if camel.Lane.Terrain.Material == camel.PreferredTerrain {
			return n * 0.3
		}

		return n * 0.15
	}

	return n

}

func (camel Dromedary) weaponModifier(n float32) interface{} {
	conditionToActivateAttribute := true

	if conditionToActivateAttribute {
		camel.eventDispatcher.Dispatch(event.WeaponEvent, camel)
	}

	return n
}

func (camel Dromedary) strengthModifier(n float32) interface{} {
	conditionWhenHitByOtherCamelsWeapons := true

	if conditionWhenHitByOtherCamelsWeapons {
		camel.eventDispatcher.Dispatch(event.StrengthEvent, camel)

		return n
	}

	return n
}

func (camel Dromedary) String() string {
	return fmt.Sprintf("Dromedary")
}
