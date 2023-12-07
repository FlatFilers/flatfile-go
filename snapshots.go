// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type CreateSnapshotRequest struct {
	// ID of sheet
	SheetId SheetId `json:"sheetId"`
	// Label for the snapshot
	Label *string `json:"label,omitempty"`
}

type GetSnapshotRequest struct {
	// Whether to include a summary in the snapshot response
	IncludeSummary bool `json:"-"`
}

type GetSnapshotRecordsRequest struct {
	// Number of records to return in a page
	PageSize *int `json:"-"`
	// Based on pageSize, which page of records to return
	PageNumber *int `json:"-"`
	// Filter records by change type
	ChangeType *ChangeType `json:"-"`
}

type ListSnapshotRequest struct {
	// ID of sheet
	SheetId SheetId `json:"-"`
}

// Snapshot ID
type SnapshotId = string

// Options to filter records in a snapshot
type ChangeType string

const (
	ChangeTypeCreatedSince ChangeType = "createdSince"
	ChangeTypeUpdatedSince ChangeType = "updatedSince"
	ChangeTypeDeletedSince ChangeType = "deletedSince"
)

func NewChangeTypeFromString(s string) (ChangeType, error) {
	switch s {
	case "createdSince":
		return ChangeTypeCreatedSince, nil
	case "updatedSince":
		return ChangeTypeUpdatedSince, nil
	case "deletedSince":
		return ChangeTypeDeletedSince, nil
	}
	var t ChangeType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ChangeType) Ptr() *ChangeType {
	return &c
}

type RestoreOptions struct {
	Created bool `json:"created"`
	Updated bool `json:"updated"`
	Deleted bool `json:"deleted"`

	_rawJSON json.RawMessage
}

func (r *RestoreOptions) UnmarshalJSON(data []byte) error {
	type unmarshaler RestoreOptions
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RestoreOptions(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RestoreOptions) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type SnapshotResponse struct {
	Data *Snapshot `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SnapshotResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SnapshotResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SnapshotResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SnapshotResponse) String() string {
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

type SnapshotsResponse struct {
	Data []*Snapshot `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SnapshotsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SnapshotsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SnapshotsResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SnapshotsResponse) String() string {
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
