// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

// Create an app
type AppCreate struct {
	Name         string      `json:"name"`
	Namespace    string      `json:"namespace"`
	Type         AppType     `json:"type,omitempty"`
	Entity       *string     `json:"entity,omitempty"`
	EntityPlural *string     `json:"entityPlural,omitempty"`
	Icon         *string     `json:"icon,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AppCreate) UnmarshalJSON(data []byte) error {
	type unmarshaler AppCreate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppCreate(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AppCreate) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// Update an app
type AppPatch struct {
	Name         *string     `json:"name,omitempty"`
	Namespace    *string     `json:"namespace,omitempty"`
	Entity       *string     `json:"entity,omitempty"`
	EntityPlural *string     `json:"entityPlural,omitempty"`
	Icon         *string     `json:"icon,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`
	ActivatedAt  *time.Time  `json:"activatedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AppPatch) UnmarshalJSON(data []byte) error {
	type unmarshaler AppPatch
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppPatch(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AppPatch) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AppResponse struct {
	Data *App `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AppResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AppResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AppResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AppsResponse struct {
	Data []*App `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AppsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AppsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppsResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AppsResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}
