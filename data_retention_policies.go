// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type ListDataRetentionPoliciesRequest struct {
	// The associated Environment ID of the policy.
	EnvironmentId *EnvironmentId `json:"-" url:"environmentId,omitempty"`
}

// Data Retention Policy ID
type DataRetentionPolicyId = string

// A data retention policy belonging to an environment
type DataRetentionPolicy struct {
	Type          DataRetentionPolicyEnum `json:"type" url:"type"`
	Period        int                     `json:"period" url:"period"`
	EnvironmentId EnvironmentId           `json:"environmentId" url:"environmentId"`
	Id            DataRetentionPolicyId   `json:"id" url:"id"`
	// Date the policy was created
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// Date the policy was last updated
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DataRetentionPolicy) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DataRetentionPolicy) UnmarshalJSON(data []byte) error {
	type embed DataRetentionPolicy
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = DataRetentionPolicy(unmarshaler.embed)
	d.CreatedAt = unmarshaler.CreatedAt.Time()
	d.UpdatedAt = unmarshaler.UpdatedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DataRetentionPolicy) MarshalJSON() ([]byte, error) {
	type embed DataRetentionPolicy
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*d),
		CreatedAt: core.NewDateTime(d.CreatedAt),
		UpdatedAt: core.NewDateTime(d.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (d *DataRetentionPolicy) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DataRetentionPolicyConfig struct {
	Type          DataRetentionPolicyEnum `json:"type" url:"type"`
	Period        int                     `json:"period" url:"period"`
	EnvironmentId EnvironmentId           `json:"environmentId" url:"environmentId"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DataRetentionPolicyConfig) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DataRetentionPolicyConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler DataRetentionPolicyConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DataRetentionPolicyConfig(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DataRetentionPolicyConfig) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// The type of data retention policy on an environment
type DataRetentionPolicyEnum string

const (
	DataRetentionPolicyEnumLastActivity DataRetentionPolicyEnum = "lastActivity"
	DataRetentionPolicyEnumSinceCreated DataRetentionPolicyEnum = "sinceCreated"
)

func NewDataRetentionPolicyEnumFromString(s string) (DataRetentionPolicyEnum, error) {
	switch s {
	case "lastActivity":
		return DataRetentionPolicyEnumLastActivity, nil
	case "sinceCreated":
		return DataRetentionPolicyEnumSinceCreated, nil
	}
	var t DataRetentionPolicyEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DataRetentionPolicyEnum) Ptr() *DataRetentionPolicyEnum {
	return &d
}

type DataRetentionPolicyResponse struct {
	Data *DataRetentionPolicy `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DataRetentionPolicyResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DataRetentionPolicyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DataRetentionPolicyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DataRetentionPolicyResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DataRetentionPolicyResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type ListDataRetentionPoliciesResponse struct {
	Data []*DataRetentionPolicy `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListDataRetentionPoliciesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListDataRetentionPoliciesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListDataRetentionPoliciesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListDataRetentionPoliciesResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListDataRetentionPoliciesResponse) String() string {
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
