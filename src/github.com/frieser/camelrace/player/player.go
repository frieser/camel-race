package player

import (
	"github.com/frieser/camelrace/camel"
)

//Player type with the transport an basic info
type Player struct {
	Name      string
	Transport camel.Movable
	Distance  float32
}
