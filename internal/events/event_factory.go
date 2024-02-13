package events

// Event interface
type Event interface {
	call() (string, error)
}

// EventFactory struct
type EventFactory struct {
	events map[EventName]Event
}

func (ef *EventFactory) getEvent(name EventName) Event {
	return ef.events[name]
}

func (ef *EventFactory) registerEvent(name EventName, event Event) {
	ef.events[name] = event
}

// Call method for EventFactory
func (ef *EventFactory) Call(event EventName) (string, error) {
	return ef.getEvent(event).call()
}

// NewEventFactory function
func NewEventFactory() *EventFactory {
	ef := &EventFactory{
		events: make(map[EventName]Event),
	}
	ef.registerEvent(LOGOUT_USER, newLogoutUser())
	return ef
}

// EventName type
type EventName int

const (
	LOGOUT_USER EventName = iota
)
