// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type ListViewsRequest struct {
	// The associated sheet ID of the view.
	SheetId SheetId `json:"-" url:"sheetId"`
	// Number of prompts to return in a page (default 10)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of prompts to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
}

// View ID
type ViewId = string

type ListViewsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty" url:"pagination,omitempty"`
	Data       []*View     `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListViewsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListViewsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListViewsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListViewsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListViewsResponse) String() string {
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

// A view
type View struct {
	// The ID of the view
	Id ViewId `json:"id" url:"id"`
	// The associated sheet ID of the view
	SheetId SheetId `json:"sheetId" url:"sheetId"`
	// The name of the view
	Name string `json:"name" url:"name"`
	// The view filters of the view
	Config *ViewConfig `json:"config,omitempty" url:"config,omitempty"`
	// ID of the actor who created the view
	CreatedBy string `json:"createdBy" url:"createdBy"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *View) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *View) UnmarshalJSON(data []byte) error {
	type unmarshaler View
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = View(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *View) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

// The configuration of a view. Filters, sorting, and search query.
type ViewConfig struct {
	// Deprecated, use `commitId` instead.
	VersionId *VersionId `json:"versionId,omitempty" url:"versionId,omitempty"`
	CommitId  *CommitId  `json:"commitId,omitempty" url:"commitId,omitempty"`
	// Deprecated, use `sinceCommitId` instead.
	SinceVersionId *VersionId     `json:"sinceVersionId,omitempty" url:"sinceVersionId,omitempty"`
	SinceCommitId  *CommitId      `json:"sinceCommitId,omitempty" url:"sinceCommitId,omitempty"`
	SortField      *SortField     `json:"sortField,omitempty" url:"sortField,omitempty"`
	SortDirection  *SortDirection `json:"sortDirection,omitempty" url:"sortDirection,omitempty"`
	Filter         *Filter        `json:"filter,omitempty" url:"filter,omitempty"`
	// Name of field by which to filter records
	FilterField *FilterField `json:"filterField,omitempty" url:"filterField,omitempty"`
	SearchValue *SearchValue `json:"searchValue,omitempty" url:"searchValue,omitempty"`
	SearchField *SearchField `json:"searchField,omitempty" url:"searchField,omitempty"`
	// The Record Ids param (ids) is a list of record ids that can be passed to several record endpoints allowing the user to identify specific records to INCLUDE in the query, or specific records to EXCLUDE, depending on whether or not filters are being applied. When passing a query param that filters the record dataset, such as 'searchValue', or a 'filter' of 'valid' | 'error' | 'all', the 'ids' param will EXCLUDE those records from the filtered results. For basic queries that do not filter the dataset, passing record ids in the 'ids' param will limit the dataset to INCLUDE just those specific records. Maximum of 100 allowed.
	Ids []RecordId `json:"ids,omitempty" url:"ids,omitempty"`
	// Number of records to return in a page (default 10,000)
	PageSize *int `json:"pageSize,omitempty" url:"pageSize,omitempty"`
	// Based on pageSize, which page of records to return (Note - numbers start at 1)
	PageNumber *int `json:"pageNumber,omitempty" url:"pageNumber,omitempty"`
	// **DEPRECATED** Use GET /sheets/:sheetId/counts
	IncludeCounts *bool `json:"includeCounts,omitempty" url:"includeCounts,omitempty"`
	// The length of the record result set, returned as counts.total
	IncludeLength *bool `json:"includeLength,omitempty" url:"includeLength,omitempty"`
	// If true, linked records will be included in the results. Defaults to false.
	IncludeLinks *bool `json:"includeLinks,omitempty" url:"includeLinks,omitempty"`
	// Include error messages, defaults to false.
	IncludeMessages *bool `json:"includeMessages,omitempty" url:"includeMessages,omitempty"`
	// if "for" is provided, the query parameters will be pulled from the event payload
	For *EventId `json:"for,omitempty" url:"for,omitempty"`
	// An FFQL query used to filter the result set
	Q *string `json:"q,omitempty" url:"q,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *ViewConfig) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *ViewConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ViewConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = ViewConfig(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *ViewConfig) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type ViewCreate struct {
	SheetId SheetId     `json:"sheetId" url:"sheetId"`
	Name    string      `json:"name" url:"name"`
	Config  *ViewConfig `json:"config,omitempty" url:"config,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *ViewCreate) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *ViewCreate) UnmarshalJSON(data []byte) error {
	type unmarshaler ViewCreate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = ViewCreate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *ViewCreate) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type ViewResponse struct {
	Data *View `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *ViewResponse) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *ViewResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ViewResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = ViewResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *ViewResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type ViewUpdate struct {
	Name   *string     `json:"name,omitempty" url:"name,omitempty"`
	Config *ViewConfig `json:"config,omitempty" url:"config,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *ViewUpdate) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *ViewUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler ViewUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = ViewUpdate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *ViewUpdate) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}
