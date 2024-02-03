package events

// Event interface
type Event interface {
	Call() (string, error)
	GetName() string
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

func GetEventFactory() *EventFactory {
	ef := &EventFactory{
		events: make(map[string]Event),
	}
	ef.RegisterEvent("logout_user", getLogoutUser())
	return ef
}
