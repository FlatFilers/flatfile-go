// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
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
	// Returns records that were changed in that version in that version and only those records.
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

type Property struct {
	Type       string
	String     *StringProperty
	Number     *NumberProperty
	Boolean    *BooleanProperty
	Date       *DateProperty
	Enum       *EnumProperty
	Reference  *ReferenceProperty
	StringList *StringListProperty
	EnumList   *EnumListProperty
}

func NewPropertyFromString(value *StringProperty) *Property {
	return &Property{Type: "string", String: value}
}

func NewPropertyFromNumber(value *NumberProperty) *Property {
	return &Property{Type: "number", Number: value}
}

func NewPropertyFromBoolean(value *BooleanProperty) *Property {
	return &Property{Type: "boolean", Boolean: value}
}

func NewPropertyFromDate(value *DateProperty) *Property {
	return &Property{Type: "date", Date: value}
}

func NewPropertyFromEnum(value *EnumProperty) *Property {
	return &Property{Type: "enum", Enum: value}
}

func NewPropertyFromReference(value *ReferenceProperty) *Property {
	return &Property{Type: "reference", Reference: value}
}

func NewPropertyFromStringList(value *StringListProperty) *Property {
	return &Property{Type: "string-list", StringList: value}
}

func NewPropertyFromEnumList(value *EnumListProperty) *Property {
	return &Property{Type: "enum-list", EnumList: value}
}

func (p *Property) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "string":
		value := new(StringProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.String = value
	case "number":
		value := new(NumberProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Number = value
	case "boolean":
		value := new(BooleanProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Boolean = value
	case "date":
		value := new(DateProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Date = value
	case "enum":
		value := new(EnumProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Enum = value
	case "reference":
		value := new(ReferenceProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Reference = value
	case "string-list":
		value := new(StringListProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.StringList = value
	case "enum-list":
		value := new(EnumListProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.EnumList = value
	}
	return nil
}

func (p Property) MarshalJSON() ([]byte, error) {
	switch p.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", p.Type, p)
	case "string":
		var marshaler = struct {
			Type string `json:"type"`
			*StringProperty
		}{
			Type:           p.Type,
			StringProperty: p.String,
		}
		return json.Marshal(marshaler)
	case "number":
		var marshaler = struct {
			Type string `json:"type"`
			*NumberProperty
		}{
			Type:           p.Type,
			NumberProperty: p.Number,
		}
		return json.Marshal(marshaler)
	case "boolean":
		var marshaler = struct {
			Type string `json:"type"`
			*BooleanProperty
		}{
			Type:            p.Type,
			BooleanProperty: p.Boolean,
		}
		return json.Marshal(marshaler)
	case "date":
		var marshaler = struct {
			Type string `json:"type"`
			*DateProperty
		}{
			Type:         p.Type,
			DateProperty: p.Date,
		}
		return json.Marshal(marshaler)
	case "enum":
		var marshaler = struct {
			Type string `json:"type"`
			*EnumProperty
		}{
			Type:         p.Type,
			EnumProperty: p.Enum,
		}
		return json.Marshal(marshaler)
	case "reference":
		var marshaler = struct {
			Type string `json:"type"`
			*ReferenceProperty
		}{
			Type:              p.Type,
			ReferenceProperty: p.Reference,
		}
		return json.Marshal(marshaler)
	case "string-list":
		var marshaler = struct {
			Type string `json:"type"`
			*StringListProperty
		}{
			Type:               p.Type,
			StringListProperty: p.StringList,
		}
		return json.Marshal(marshaler)
	case "enum-list":
		var marshaler = struct {
			Type string `json:"type"`
			*EnumListProperty
		}{
			Type:             p.Type,
			EnumListProperty: p.EnumList,
		}
		return json.Marshal(marshaler)
	}
}

type PropertyVisitor interface {
	VisitString(*StringProperty) error
	VisitNumber(*NumberProperty) error
	VisitBoolean(*BooleanProperty) error
	VisitDate(*DateProperty) error
	VisitEnum(*EnumProperty) error
	VisitReference(*ReferenceProperty) error
	VisitStringList(*StringListProperty) error
	VisitEnumList(*EnumListProperty) error
}

func (p *Property) Accept(visitor PropertyVisitor) error {
	switch p.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", p.Type, p)
	case "string":
		return visitor.VisitString(p.String)
	case "number":
		return visitor.VisitNumber(p.Number)
	case "boolean":
		return visitor.VisitBoolean(p.Boolean)
	case "date":
		return visitor.VisitDate(p.Date)
	case "enum":
		return visitor.VisitEnum(p.Enum)
	case "reference":
		return visitor.VisitReference(p.Reference)
	case "string-list":
		return visitor.VisitStringList(p.StringList)
	case "enum-list":
		return visitor.VisitEnumList(p.EnumList)
	}
}

type CellsResponse struct {
	Data CellsResponseData `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CellsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CellsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CellsResponse(value)
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

// When true, excludes duplicate values
type Distinct = bool

// Returns results from the given field only. Otherwise all field cells are returned
type FieldKey = string

// When both distinct and includeCounts are true, the count of distinct field values will be returned
type IncludeCounts = bool

type RecordCountsResponse struct {
	Data *RecordCountsResponseData `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RecordCountsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RecordCountsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RecordCountsResponse(value)
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

type SheetResponse struct {
	Data *Sheet `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SheetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SheetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SheetResponse(value)
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
