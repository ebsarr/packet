package extpackngo

import (
	"fmt"

	"github.com/packethost/packngo"
)

const eventBasePath = "/events"

// EventService interface defines available event methods
type EventService interface {
	List(deviceID string) ([]Event, *Response, error)
	Get(string) (*Event, *Response, error)
}

type eventsRoot struct {
	Events []Event `json:"events"`
}

// Event represents a Packet Event
type Event struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Body          string              `json:"body"`
	Create        string              `json:"created_at"`
	Interpolated  string              `json:"interpolated"`
	Href          string              `json:"href"`
	Relationships []map[string]string `json:"relationships"`
}

func (e Event) String() string {
	return packngo.Stringify(e)
}

// EventServiceOp implements EventService
type EventServiceOp struct {
	client *Client
}

// List returns Events of a device
func (e *EventServiceOp) List(deviceID string) ([]Event, *Response, error) {
	path := fmt.Sprintf("devices/%s/%s", deviceID, eventBasePath)

	req, err := e.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(eventsRoot)
	resp, err := e.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Events, resp, err
}

// Get returns and event by ID
func (e *EventServiceOp) Get(eventID string) (*Event, *Response, error) {
	path := fmt.Sprintf("%s/%s", eventBasePath, eventID)

	req, err := e.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	event := new(Event)
	resp, err := e.client.Do(req, event)
	if err != nil {
		return nil, resp, err
	}

	return event, resp, err
}
