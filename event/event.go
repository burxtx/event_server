package event

import (
	"git.n.xiaomi.com/op-basic/event_server/event/utils"
)

type HostID string

// Event is the central of event domain model
// play as aggragate root
type Event struct {
	ID        HostID
	EventType EventType
}

// NewUser creates a new user
func NewEvent(eventType EventType) *Event {
	return &Event{
		ID:        HostID(utils.NewHostID()),
		EventType: eventType,
	}
}

// Repository provides access to a user store.
type EventRepository interface {
	OnCreated() error
	Publish() error
	Subscribe() error
}
