package events

// Event interface
type Event interface {
	Call() (string, error)
}

// EventFactory struct
type EventFactory struct {
	events map[string]Event
}

// GetEvent method for EventFactory
func (ef *EventFactory) GetEvent(name string) Event {
	return ef.events[name]
}

// RegisterEvent method for EventFactory
func (ef *EventFactory) RegisterEvent(name string, event Event) {
	ef.events[name] = event
}

// GetEventFactory function
func GetEventFactory() *EventFactory {
	ef := &EventFactory{
		events: make(map[string]Event),
	}
	ef.RegisterEvent("logout_user", getLogoutUser())
	return ef
}
