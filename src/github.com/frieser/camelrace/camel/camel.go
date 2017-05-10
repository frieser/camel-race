package camel

import (
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
)

type camel struct {
	Lane             lane.Lane
	Attributes       []Attribute
	PreferredTerrain string
	eventDispatcher  event.EventDispatcher
}
