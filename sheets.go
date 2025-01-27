// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type GetFieldValuesRequest struct {
	FieldKey      *FieldKey      `json:"-" url:"fieldKey,omitempty"`
	SortField     *SortField     `json:"-" url:"sortField,omitempty"`
	SortDirection *SortDirection `json:"-" url:"sortDirection,omitempty"`
	Filter        *Filter        `json:"-" url:"filter,omitempty"`
	// Name of field by which to filter records
	FilterField *FilterField `json:"-" url:"filterField,omitempty"`
	// Number of records to return in a page (default 1000 if pageNumber included)
	PageSize *PageSize `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of records to return
	PageNumber *PageNumber `json:"-" url:"pageNumber,omitempty"`
	// Must be set to true
	Distinct      Distinct       `json:"-" url:"distinct"`
	IncludeCounts *IncludeCounts `json:"-" url:"includeCounts,omitempty"`
	// A value to find for a given field in a sheet. Wrap the value in "" for exact match
	SearchValue *SearchValue `json:"-" url:"searchValue,omitempty"`
}

type GetRecordCountsRequest struct {
	// Returns records that were changed in that version and only those records.
	VersionId *string `json:"-" url:"versionId,omitempty"`
	// Deprecated, use `sinceCommitId` instead.
	SinceVersionId *VersionId `json:"-" url:"sinceVersionId,omitempty"`
	// Returns records that were changed in that version in addition to any records from versions after that version.
	CommitId *CommitId `json:"-" url:"commitId,omitempty"`
	// Listing a commit ID here will return all records since the specified commit.
	SinceCommitId *CommitId `json:"-" url:"sinceCommitId,omitempty"`
	// Options to filter records
	Filter *Filter `json:"-" url:"filter,omitempty"`
	// The field to filter the data on.
	FilterField *FilterField `json:"-" url:"filterField,omitempty"`
	// The value to search for data on.
	SearchValue *SearchValue `json:"-" url:"searchValue,omitempty"`
	// The field to search for data on.
	SearchField *SearchField `json:"-" url:"searchField,omitempty"`
	// If true, the counts for each field will also be returned
	ByField *bool `json:"-" url:"byField,omitempty"`
	// An FFQL query used to filter the result set to be counted
	Q *string `json:"-" url:"q,omitempty"`
}

type GetRecordsCsvRequest struct {
	// Deprecated, use `sinceCommitId` instead.
	VersionId *string `json:"-" url:"versionId,omitempty"`
	// Returns records that were changed in that version  in that version and only those records.
	CommitId *CommitId `json:"-" url:"commitId,omitempty"`
	// Deprecated, use `sinceCommitId` instead.
	SinceVersionId *VersionId `json:"-" url:"sinceVersionId,omitempty"`
	// Returns records that were changed in that version in addition to any records from versions after that version.
	SinceCommitId *CommitId `json:"-" url:"sinceCommitId,omitempty"`
	// The field to sort the data on.
	SortField *SortField `json:"-" url:"sortField,omitempty"`
	// Sort direction - asc (ascending) or desc (descending)
	SortDirection *SortDirection `json:"-" url:"sortDirection,omitempty"`
	// Options to filter records
	Filter *Filter `json:"-" url:"filter,omitempty"`
	// The field to filter the data on.
	FilterField *FilterField `json:"-" url:"filterField,omitempty"`
	// The value to search for data on.
	SearchValue *SearchValue `json:"-" url:"searchValue,omitempty"`
	// The field to search for data on.
	SearchField *SearchField `json:"-" url:"searchField,omitempty"`
	// The Record Ids param (ids) is a list of record ids that can be passed to several record endpoints allowing the user to identify specific records to INCLUDE in the query, or specific records to EXCLUDE, depending on whether or not filters are being applied. When passing a query param that filters the record dataset, such as 'searchValue', or a 'filter' of 'valid' | 'error' | 'all', the 'ids' param will EXCLUDE those records from the filtered results. For basic queries that do not filter the dataset, passing record ids in the 'ids' param will limit the dataset to INCLUDE just those specific records
	Ids []*RecordId `json:"-" url:"ids,omitempty"`
}

type ListSheetCommitsRequest struct {
	// If true, only return commits that have been completed. If false, only return commits that have not been completed. If not provided, return all commits.
	Completed *bool `json:"-" url:"completed,omitempty"`
}

type ListSheetsRequest struct {
	// ID of workbook
	WorkbookId WorkbookId `json:"-" url:"workbookId"`
}

type CellValueWithCounts struct {
	Valid    *bool                `json:"valid,omitempty" url:"valid,omitempty"`
	Messages []*ValidationMessage `json:"messages,omitempty" url:"messages,omitempty"`
	// Deprecated, use record level metadata instead.
	Metadata  map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Value     *CellValueUnion        `json:"value,omitempty" url:"value,omitempty"`
	Layer     *string                `json:"layer,omitempty" url:"layer,omitempty"`
	UpdatedAt *time.Time             `json:"updatedAt,omitempty" url:"updatedAt,omitempty"`
	Counts    *RecordCounts          `json:"counts,omitempty" url:"counts,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CellValueWithCounts) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CellValueWithCounts) UnmarshalJSON(data []byte) error {
	type embed CellValueWithCounts
	var unmarshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CellValueWithCounts(unmarshaler.embed)
	c.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CellValueWithCounts) MarshalJSON() ([]byte, error) {
	type embed CellValueWithCounts
	var marshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt,omitempty"`
	}{
		embed:     embed(*c),
		UpdatedAt: core.NewOptionalDateTime(c.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (c *CellValueWithCounts) String() string {
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

type CellsResponse struct {
	Data CellsResponseData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CellsResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CellsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CellsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CellsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CellsResponse) String() string {
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

// Cell values grouped by field key
type CellsResponseData = map[string][]*CellValueWithCounts

// When true, excludes duplicate values
type Distinct = bool

// Returns results from the given field only. Otherwise all field cells are returned
type FieldKey = string

// When both distinct and includeCounts are true, the count of distinct field values will be returned
type IncludeCounts = bool

type RecordCountsResponse struct {
	Data *RecordCountsResponseData `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RecordCountsResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RecordCountsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RecordCountsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RecordCountsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RecordCountsResponse) String() string {
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

type RecordCountsResponseData struct {
	Counts  *RecordCounts `json:"counts,omitempty" url:"counts,omitempty"`
	Success bool          `json:"success" url:"success"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RecordCountsResponseData) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RecordCountsResponseData) UnmarshalJSON(data []byte) error {
	type unmarshaler RecordCountsResponseData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RecordCountsResponseData(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RecordCountsResponseData) String() string {
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

type SheetResponse struct {
	Data *Sheet `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SheetResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SheetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SheetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SheetResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SheetResponse) String() string {
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

// Changes to make to an existing sheet
type SheetUpdateRequest struct {
	// The name of the Sheet.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The slug of the Sheet.
	Slug *string `json:"slug,omitempty" url:"slug,omitempty"`
	// Useful for any contextual metadata regarding the sheet. Store any valid json
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SheetUpdateRequest) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SheetUpdateRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler SheetUpdateRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SheetUpdateRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SheetUpdateRequest) String() string {
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
