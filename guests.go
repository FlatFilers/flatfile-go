// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type GetGuestTokenRequest struct {
	// ID of space to return
	SpaceId *SpaceId `json:"-" url:"spaceId,omitempty"`
}

type ListGuestsRequest struct {
	// ID of space to return
	SpaceId SpaceId `json:"-" url:"spaceId"`
	// Email of guest to return
	Email *string `json:"-" url:"email,omitempty"`
}

// Guest ID
type GuestId = string

type CreateGuestResponse struct {
	Data []*Guest `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateGuestResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateGuestResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateGuestResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateGuestResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateGuestResponse) String() string {
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

// Configurations for the guests
type GuestConfig struct {
	EnvironmentId EnvironmentId `json:"environmentId" url:"environmentId"`
	Email         string        `json:"email" url:"email"`
	Name          string        `json:"name" url:"name"`
	Spaces        []*GuestSpace `json:"spaces,omitempty" url:"spaces,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestConfig) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestConfig(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestConfig) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Properties used to update an existing guest
type GuestConfigUpdate struct {
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty" url:"environmentId,omitempty"`
	Email         *string        `json:"email,omitempty" url:"email,omitempty"`
	Name          *string        `json:"name,omitempty" url:"name,omitempty"`
	Spaces        []*GuestSpace  `json:"spaces,omitempty" url:"spaces,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestConfigUpdate) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestConfigUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestConfigUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestConfigUpdate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestConfigUpdate) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GuestResponse struct {
	Data *Guest `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GuestTokenResponse struct {
	Data *GuestToken `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestTokenResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestTokenResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestTokenResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestTokenResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestTokenResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type Invite struct {
	GuestId GuestId `json:"guestId" url:"guestId"`
	SpaceId SpaceId `json:"spaceId" url:"spaceId"`
	// The name of the person or company sending the invitation
	FromName *string `json:"fromName,omitempty" url:"fromName,omitempty"`
	// Message to send with the invite
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *Invite) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *Invite) UnmarshalJSON(data []byte) error {
	type unmarshaler Invite
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = Invite(value)

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *Invite) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type ListGuestsResponse struct {
	Data []*Guest `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListGuestsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListGuestsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListGuestsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListGuestsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListGuestsResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}
