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
	Name               string      `json:"name" url:"name"`
	Namespace          string      `json:"namespace" url:"namespace"`
	Type               AppType     `json:"type" url:"type"`
	Entity             *string     `json:"entity,omitempty" url:"entity,omitempty"`
	EntityPlural       *string     `json:"entityPlural,omitempty" url:"entityPlural,omitempty"`
	Icon               *string     `json:"icon,omitempty" url:"icon,omitempty"`
	Metadata           interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	EnvironmentFilters interface{} `json:"environmentFilters,omitempty" url:"environmentFilters,omitempty"`
	Blueprint          interface{} `json:"blueprint,omitempty" url:"blueprint,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AppCreate) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AppCreate) UnmarshalJSON(data []byte) error {
	type unmarshaler AppCreate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppCreate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

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
	Name               *string     `json:"name,omitempty" url:"name,omitempty"`
	Namespace          *string     `json:"namespace,omitempty" url:"namespace,omitempty"`
	Entity             *string     `json:"entity,omitempty" url:"entity,omitempty"`
	EntityPlural       *string     `json:"entityPlural,omitempty" url:"entityPlural,omitempty"`
	Icon               *string     `json:"icon,omitempty" url:"icon,omitempty"`
	Metadata           interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	EnvironmentFilters interface{} `json:"environmentFilters,omitempty" url:"environmentFilters,omitempty"`
	Blueprint          interface{} `json:"blueprint,omitempty" url:"blueprint,omitempty"`
	ActivatedAt        *time.Time  `json:"activatedAt,omitempty" url:"activatedAt,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AppPatch) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AppPatch) UnmarshalJSON(data []byte) error {
	type embed AppPatch
	var unmarshaler = struct {
		embed
		ActivatedAt *core.DateTime `json:"activatedAt,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = AppPatch(unmarshaler.embed)
	a.ActivatedAt = unmarshaler.ActivatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AppPatch) MarshalJSON() ([]byte, error) {
	type embed AppPatch
	var marshaler = struct {
		embed
		ActivatedAt *core.DateTime `json:"activatedAt,omitempty"`
	}{
		embed:       embed(*a),
		ActivatedAt: core.NewOptionalDateTime(a.ActivatedAt),
	}
	return json.Marshal(marshaler)
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
	Data *App `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AppResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AppResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AppResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

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
	Data []*App `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AppsResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AppsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AppsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AppsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

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

type SuccessResponse struct {
	Success bool `json:"success" url:"success"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SuccessResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SuccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SuccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SuccessResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SuccessResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
