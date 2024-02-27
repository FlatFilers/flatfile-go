// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type GetEventTokenRequest struct {
	// The resource ID of the event stream (space or environment id)
	Scope *string `json:"-" url:"scope,omitempty"`
	// The space ID of the event stream
	SpaceId *SpaceId `json:"-" url:"spaceId,omitempty"`
}

type ListEventsRequest struct {
	// Filter by environment
	EnvironmentId *EnvironmentId `json:"-" url:"environmentId,omitempty"`
	// Filter by space
	SpaceId *SpaceId `json:"-" url:"spaceId,omitempty"`
	// Filter by event domain
	Domain *string `json:"-" url:"domain,omitempty"`
	// Filter by event topic
	Topic *string `json:"-" url:"topic,omitempty"`
	// Filter by event timestamp
	Since *time.Time `json:"-" url:"since,omitempty"`
	// Number of results to return in a page (default 10)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of results to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
	// Include acknowledged events
	IncludeAcknowledged *bool `json:"-" url:"includeAcknowledged,omitempty"`
}

// Properties used to create a new event
type CreateEventConfig struct {
	// The domain of the event
	Domain Domain `json:"domain,omitempty" url:"domain,omitempty"`
	// The context of the event
	Context *Context `json:"context,omitempty" url:"context,omitempty"`
	// The attributes of the event
	Attributes *EventAttributes `json:"attributes,omitempty" url:"attributes,omitempty"`
	// The callback url to acknowledge the event
	CallbackUrl *string `json:"callbackUrl,omitempty" url:"callbackUrl,omitempty"`
	// The url to retrieve the data associated with the event
	DataUrl    *string                `json:"dataUrl,omitempty" url:"dataUrl,omitempty"`
	Target     *string                `json:"target,omitempty" url:"target,omitempty"`
	Origin     *Origin                `json:"origin,omitempty" url:"origin,omitempty"`
	Namespaces []string               `json:"namespaces,omitempty" url:"namespaces,omitempty"`
	Topic      EventTopic             `json:"topic,omitempty" url:"topic,omitempty"`
	Payload    map[string]interface{} `json:"payload,omitempty" url:"payload,omitempty"`
	// Date the event was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty" url:"deletedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CreateEventConfig) UnmarshalJSON(data []byte) error {
	type embed CreateEventConfig
	var unmarshaler = struct {
		embed
		DeletedAt *core.DateTime `json:"deletedAt,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateEventConfig(unmarshaler.embed)
	c.DeletedAt = unmarshaler.DeletedAt.TimePtr()
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateEventConfig) MarshalJSON() ([]byte, error) {
	type embed CreateEventConfig
	var marshaler = struct {
		embed
		DeletedAt *core.DateTime `json:"deletedAt,omitempty"`
	}{
		embed:     embed(*c),
		DeletedAt: core.NewOptionalDateTime(c.DeletedAt),
	}
	return json.Marshal(marshaler)
}

func (c *CreateEventConfig) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}
