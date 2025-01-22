// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
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

type Guest struct {
	EnvironmentId EnvironmentId `json:"environmentId" url:"environmentId"`
	Email         string        `json:"email" url:"email"`
	Name          string        `json:"name" url:"name"`
	Spaces        []*GuestSpace `json:"spaces,omitempty" url:"spaces,omitempty"`
	Id            GuestId       `json:"id" url:"id"`
	// Date the guest object was created
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// Date the guest object was last updated
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *Guest) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *Guest) UnmarshalJSON(data []byte) error {
	type embed Guest
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = Guest(unmarshaler.embed)
	g.CreatedAt = unmarshaler.CreatedAt.Time()
	g.UpdatedAt = unmarshaler.UpdatedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *Guest) MarshalJSON() ([]byte, error) {
	type embed Guest
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*g),
		CreatedAt: core.NewDateTime(g.CreatedAt),
		UpdatedAt: core.NewDateTime(g.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (g *Guest) String() string {
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

type GuestSpace struct {
	Id           SpaceId          `json:"id" url:"id"`
	Workbooks    []*GuestWorkbook `json:"workbooks,omitempty" url:"workbooks,omitempty"`
	LastAccessed *time.Time       `json:"lastAccessed,omitempty" url:"lastAccessed,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestSpace) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestSpace) UnmarshalJSON(data []byte) error {
	type embed GuestSpace
	var unmarshaler = struct {
		embed
		LastAccessed *core.DateTime `json:"lastAccessed,omitempty"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GuestSpace(unmarshaler.embed)
	g.LastAccessed = unmarshaler.LastAccessed.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestSpace) MarshalJSON() ([]byte, error) {
	type embed GuestSpace
	var marshaler = struct {
		embed
		LastAccessed *core.DateTime `json:"lastAccessed,omitempty"`
	}{
		embed:        embed(*g),
		LastAccessed: core.NewOptionalDateTime(g.LastAccessed),
	}
	return json.Marshal(marshaler)
}

func (g *GuestSpace) String() string {
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

type GuestToken struct {
	// The token used to authenticate the guest
	Token string `json:"token" url:"token"`
	Valid bool   `json:"valid" url:"valid"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestToken) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestToken) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestToken
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestToken(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestToken) String() string {
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

type GuestWorkbook struct {
	Id WorkbookId `json:"id" url:"id"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GuestWorkbook) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GuestWorkbook) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestWorkbook
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestWorkbook(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestWorkbook) String() string {
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
