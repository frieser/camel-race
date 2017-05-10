package event

import (
	"reflect"
)

const WeaponEvent string = "WEAPON"
const TurboBoostEvent string = "TURBOBOOST"
const StrengthEvent string = "STRENGTH"

type Event interface {
	Dispatch(args ...interface{}) error
	On(f func(args ...interface{})) error
}

type event struct {
	listeners []func(args ...interface{})
}

func New() *event {
	return &event{}
}

//Trigger an event in the map of events
func (event *event) Dispatch(args ...interface{}) error {
	arguments := make([]reflect.Value, 0, len(args))
	argTypes := make([]reflect.Type, 0, len(args))

	for _, v := range args {
		arguments = append(arguments, reflect.ValueOf(v))
		argTypes = append(argTypes, reflect.TypeOf(v))
	}

	for _, fn := range event.listeners {
		fn(arguments)
	}

	return nil
}

//Chain the execution of a function when a certain event is triggered
func (event *event) On(f func(args ...interface{})) error {
	event.listeners = append(event.listeners, f)

	return nil
}
