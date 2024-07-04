// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type GetEnvironmentEventTokenRequest struct {
	// ID of environment to return
	EnvironmentId EnvironmentId `json:"-" url:"environmentId"`
}

type ListEnvironmentsRequest struct {
	// Number of environments to return in a page (default 10)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of environments to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
}

type Environment struct {
	Id        EnvironmentId `json:"id" url:"id"`
	AccountId AccountId     `json:"accountId" url:"accountId"`
	// The name of the environment
	Name string `json:"name" url:"name"`
	// Whether or not the environment is a production environment
	IsProd              bool                      `json:"isProd" url:"isProd"`
	GuestAuthentication []GuestAuthenticationEnum `json:"guestAuthentication,omitempty" url:"guestAuthentication,omitempty"`
	Features            map[string]interface{}    `json:"features,omitempty" url:"features,omitempty"`
	Metadata            map[string]interface{}    `json:"metadata,omitempty" url:"metadata,omitempty"`
	TranslationsPath    *string                   `json:"translationsPath,omitempty" url:"translationsPath,omitempty"`
	Namespaces          []string                  `json:"namespaces,omitempty" url:"namespaces,omitempty"`
	LanguageOverride    *string                   `json:"languageOverride,omitempty" url:"languageOverride,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *Environment) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *Environment) UnmarshalJSON(data []byte) error {
	type unmarshaler Environment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = Environment(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *Environment) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Properties used to create a new environment
type EnvironmentConfigCreate struct {
	// The name of the environment
	Name string `json:"name" url:"name"`
	// Whether or not the environment is a production environment
	IsProd              bool                      `json:"isProd" url:"isProd"`
	GuestAuthentication []GuestAuthenticationEnum `json:"guestAuthentication,omitempty" url:"guestAuthentication,omitempty"`
	Metadata            map[string]interface{}    `json:"metadata,omitempty" url:"metadata,omitempty"`
	TranslationsPath    *string                   `json:"translationsPath,omitempty" url:"translationsPath,omitempty"`
	Namespaces          []string                  `json:"namespaces,omitempty" url:"namespaces,omitempty"`
	LanguageOverride    *string                   `json:"languageOverride,omitempty" url:"languageOverride,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EnvironmentConfigCreate) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EnvironmentConfigCreate) UnmarshalJSON(data []byte) error {
	type unmarshaler EnvironmentConfigCreate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EnvironmentConfigCreate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EnvironmentConfigCreate) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Properties used to update an environment
type EnvironmentConfigUpdate struct {
	// The name of the environment
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// Whether or not the environment is a production environment
	IsProd              *bool                     `json:"isProd,omitempty" url:"isProd,omitempty"`
	GuestAuthentication []GuestAuthenticationEnum `json:"guestAuthentication,omitempty" url:"guestAuthentication,omitempty"`
	Metadata            map[string]interface{}    `json:"metadata,omitempty" url:"metadata,omitempty"`
	TranslationsPath    *string                   `json:"translationsPath,omitempty" url:"translationsPath,omitempty"`
	Namespaces          []string                  `json:"namespaces,omitempty" url:"namespaces,omitempty"`
	LanguageOverride    *string                   `json:"languageOverride,omitempty" url:"languageOverride,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EnvironmentConfigUpdate) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EnvironmentConfigUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler EnvironmentConfigUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EnvironmentConfigUpdate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EnvironmentConfigUpdate) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EnvironmentResponse struct {
	Data *Environment `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EnvironmentResponse) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EnvironmentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EnvironmentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EnvironmentResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EnvironmentResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ListEnvironmentsResponse struct {
	Data       []*Environment `json:"data,omitempty" url:"data,omitempty"`
	Pagination *Pagination    `json:"pagination,omitempty" url:"pagination,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListEnvironmentsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListEnvironmentsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListEnvironmentsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListEnvironmentsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListEnvironmentsResponse) String() string {
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
