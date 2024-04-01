// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

// Properties used to update an account
type AccountPatch struct {
	DefaultAppId AppId `json:"defaultAppId" url:"defaultAppId"`

	_rawJSON json.RawMessage
}

func (a *AccountPatch) UnmarshalJSON(data []byte) error {
	type unmarshaler AccountPatch
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccountPatch(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccountPatch) String() string {
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

type AccountResponse struct {
	Data *Account `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AccountResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccountResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccountResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccountResponse) String() string {
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
