package event

import "errors"

type EventDispatcher struct {
	events map[string]Event
}

func NewEventsMap() EventDispatcher {
	return EventDispatcher{events: make(map[string]Event)}
}

//Trigger an event in the map of events
func (eventMap *EventDispatcher) Dispatch(name string, args ...interface{}) error {
	event, ok := eventMap.events[name]
	if !ok {
		return errors.New("Event not found")
	}

	return event.Dispatch(args)
}

//Chain the execution of a function when a certain event is triggered
func (eventMap *EventDispatcher) On(name string, f func(args ...interface{})) error {
	event, ok := eventMap.events[name]

	if !ok {
		event = New()
		eventMap.events[name] = event
	}
	return event.On(f)
}
