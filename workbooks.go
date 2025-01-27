// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type ListWorkbookCommitsRequest struct {
	// If true, only return commits that have been completed. If false, only return commits that have not been completed. If not provided, return all commits.
	Completed *bool `json:"-" url:"completed,omitempty"`
}

type ListWorkbooksRequest struct {
	// The associated Space ID of the Workbook.
	SpaceId *SpaceId `json:"-" url:"spaceId,omitempty"`
	// Filter by name. Precede with - to negate the filter
	Name *string `json:"-" url:"name,omitempty"`
	// Filter by namespace. Precede with - to negate the filter
	Namespace *string `json:"-" url:"namespace,omitempty"`
	// Filter by label. Precede with - to negate the filter
	Label *string `json:"-" url:"label,omitempty"`
	// Filter by treatment.
	Treatment *string `json:"-" url:"treatment,omitempty"`
	// Include sheets for the workbook (default true)
	IncludeSheets *bool `json:"-" url:"includeSheets,omitempty"`
	// Include counts for the workbook. **DEPRECATED** Counts will return 0s. Use GET /sheets/:sheetId/counts
	IncludeCounts *bool `json:"-" url:"includeCounts,omitempty"`
}

type SheetConfigOrUpdate struct {
	// The name of your Sheet as it will appear to your end users.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// A sentence or two describing the purpose of your Sheet.
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// A unique identifier for your Sheet. **Required when updating a Workbook.**
	Slug *string `json:"slug,omitempty" url:"slug,omitempty"`
	// A boolean specifying whether or not this sheet is read only. Read only sheets are not editable by end users.
	Readonly *bool `json:"readonly,omitempty" url:"readonly,omitempty"`
	// Allow end users to add fields during mapping.
	AllowAdditionalFields *bool `json:"allowAdditionalFields,omitempty" url:"allowAdditionalFields,omitempty"`
	// The minimum confidence required to automatically map a field
	MappingConfidenceThreshold *float64 `json:"mappingConfidenceThreshold,omitempty" url:"mappingConfidenceThreshold,omitempty"`
	// Control Sheet-level access for all users.
	Access []SheetAccess `json:"access,omitempty" url:"access,omitempty"`
	// Where you define your Sheet’s data schema.
	Fields []*Property `json:"fields,omitempty" url:"fields,omitempty"`
	// An array of actions that end users can perform on this Sheet.
	Actions []*Action `json:"actions,omitempty" url:"actions,omitempty"`
	// The ID of the Sheet.
	Id *SheetId `json:"id,omitempty" url:"id,omitempty"`
	// The ID of the Workbook.
	WorkbookId *WorkbookId `json:"workbookId,omitempty" url:"workbookId,omitempty"`
	// Describes shape of data as well as behavior.
	Config *SheetConfig `json:"config,omitempty" url:"config,omitempty"`
	// Useful for any contextual metadata regarding the sheet. Store any valid json
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// The scoped namespace of the Sheet.
	Namespace *string `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Date the sheet was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty" url:"updatedAt,omitempty"`
	// Date the sheet was created
	CreatedAt *time.Time `json:"createdAt,omitempty" url:"createdAt,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SheetConfigOrUpdate) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SheetConfigOrUpdate) UnmarshalJSON(data []byte) error {
	type embed SheetConfigOrUpdate
	var unmarshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt,omitempty"`
		CreatedAt *core.DateTime `json:"createdAt,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SheetConfigOrUpdate(unmarshaler.embed)
	s.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	s.CreatedAt = unmarshaler.CreatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SheetConfigOrUpdate) MarshalJSON() ([]byte, error) {
	type embed SheetConfigOrUpdate
	var marshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt,omitempty"`
		CreatedAt *core.DateTime `json:"createdAt,omitempty"`
	}{
		embed:     embed(*s),
		UpdatedAt: core.NewOptionalDateTime(s.UpdatedAt),
		CreatedAt: core.NewOptionalDateTime(s.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (s *SheetConfigOrUpdate) String() string {
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

// Changes to make to an existing sheet config
type SheetConfigUpdate struct {
	// The name of your Sheet as it will appear to your end users.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// A sentence or two describing the purpose of your Sheet.
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// A unique identifier for your Sheet. **Required when updating a Workbook.**
	Slug *string `json:"slug,omitempty" url:"slug,omitempty"`
	// A boolean specifying whether or not this sheet is read only. Read only sheets are not editable by end users.
	Readonly *bool `json:"readonly,omitempty" url:"readonly,omitempty"`
	// Allow end users to add fields during mapping.
	AllowAdditionalFields *bool `json:"allowAdditionalFields,omitempty" url:"allowAdditionalFields,omitempty"`
	// The minimum confidence required to automatically map a field
	MappingConfidenceThreshold *float64 `json:"mappingConfidenceThreshold,omitempty" url:"mappingConfidenceThreshold,omitempty"`
	// Control Sheet-level access for all users.
	Access []SheetAccess `json:"access,omitempty" url:"access,omitempty"`
	// Where you define your Sheet’s data schema.
	Fields []*Property `json:"fields,omitempty" url:"fields,omitempty"`
	// An array of actions that end users can perform on this Sheet.
	Actions []*Action `json:"actions,omitempty" url:"actions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SheetConfigUpdate) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SheetConfigUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler SheetConfigUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SheetConfigUpdate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SheetConfigUpdate) String() string {
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
type SheetUpdate struct {
	// The ID of the Sheet.
	Id *SheetId `json:"id,omitempty" url:"id,omitempty"`
	// The ID of the Workbook.
	WorkbookId *WorkbookId `json:"workbookId,omitempty" url:"workbookId,omitempty"`
	// Describes shape of data as well as behavior.
	Config *SheetConfig `json:"config,omitempty" url:"config,omitempty"`
	// Useful for any contextual metadata regarding the sheet. Store any valid json
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// The scoped namespace of the Sheet.
	Namespace *string `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Date the sheet was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty" url:"updatedAt,omitempty"`
	// Date the sheet was created
	CreatedAt *time.Time `json:"createdAt,omitempty" url:"createdAt,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SheetUpdate) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SheetUpdate) UnmarshalJSON(data []byte) error {
	type embed SheetUpdate
	var unmarshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt,omitempty"`
		CreatedAt *core.DateTime `json:"createdAt,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SheetUpdate(unmarshaler.embed)
	s.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	s.CreatedAt = unmarshaler.CreatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SheetUpdate) MarshalJSON() ([]byte, error) {
	type embed SheetUpdate
	var marshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt,omitempty"`
		CreatedAt *core.DateTime `json:"createdAt,omitempty"`
	}{
		embed:     embed(*s),
		UpdatedAt: core.NewOptionalDateTime(s.UpdatedAt),
		CreatedAt: core.NewOptionalDateTime(s.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (s *SheetUpdate) String() string {
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

// Properties used to create a new Workbook
type CreateWorkbookConfig struct {
	// The name of the Workbook.
	Name string `json:"name" url:"name"`
	// An optional list of labels for the Workbook.
	Labels []string `json:"labels,omitempty" url:"labels,omitempty"`
	// Space to associate with the Workbook.
	SpaceId *SpaceId `json:"spaceId,omitempty" url:"spaceId,omitempty"`
	// Environment to associate with the Workbook
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty" url:"environmentId,omitempty"`
	// Optional namespace to apply to the Workbook.
	Namespace *string `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Sheets to create on the Workbook.
	Sheets []*SheetConfig `json:"sheets,omitempty" url:"sheets,omitempty"`
	// Actions to create on the Workbook.
	Actions []*Action `json:"actions,omitempty" url:"actions,omitempty"`
	// The Workbook settings.
	Settings *WorkbookConfigSettings `json:"settings,omitempty" url:"settings,omitempty"`
	// Metadata for the workbook
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// Treatments for the workbook
	Treatments []WorkbookTreatments `json:"treatments,omitempty" url:"treatments,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateWorkbookConfig) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateWorkbookConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateWorkbookConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateWorkbookConfig(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateWorkbookConfig) String() string {
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

type ListWorkbooksResponse struct {
	Data []*Workbook `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListWorkbooksResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListWorkbooksResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListWorkbooksResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListWorkbooksResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListWorkbooksResponse) String() string {
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

// A collection of one or more sheets
type Workbook struct {
	// ID of the Workbook.
	Id WorkbookId `json:"id" url:"id"`
	// Name of the Workbook.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// Associated Space ID of the Workbook.
	SpaceId SpaceId `json:"spaceId" url:"spaceId"`
	// Associated Environment ID of the Workbook.
	EnvironmentId EnvironmentId `json:"environmentId" url:"environmentId"`
	// A list of Sheets associated with the Workbook.
	Sheets []*Sheet `json:"sheets,omitempty" url:"sheets,omitempty"`
	// A list of labels for the Workbook.
	Labels []string `json:"labels,omitempty" url:"labels,omitempty"`
	// A list of Actions associated with the Workbook.
	Actions []*Action `json:"actions,omitempty" url:"actions,omitempty"`
	// The Workbook settings.
	Settings *WorkbookConfigSettings `json:"settings,omitempty" url:"settings,omitempty"`
	// Metadata for the workbook
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// Treatments for the workbook
	Treatments []WorkbookTreatments `json:"treatments,omitempty" url:"treatments,omitempty"`
	Namespace  *string              `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Date the workbook was last updated
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`
	// Date the workbook was created
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// Date the workbook was created
	ExpiredAt *time.Time `json:"expiredAt,omitempty" url:"expiredAt,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (w *Workbook) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *Workbook) UnmarshalJSON(data []byte) error {
	type embed Workbook
	var unmarshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt"`
		CreatedAt *core.DateTime `json:"createdAt"`
		ExpiredAt *core.DateTime `json:"expiredAt,omitempty"`
	}{
		embed: embed(*w),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*w = Workbook(unmarshaler.embed)
	w.UpdatedAt = unmarshaler.UpdatedAt.Time()
	w.CreatedAt = unmarshaler.CreatedAt.Time()
	w.ExpiredAt = unmarshaler.ExpiredAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties

	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *Workbook) MarshalJSON() ([]byte, error) {
	type embed Workbook
	var marshaler = struct {
		embed
		UpdatedAt *core.DateTime `json:"updatedAt"`
		CreatedAt *core.DateTime `json:"createdAt"`
		ExpiredAt *core.DateTime `json:"expiredAt,omitempty"`
	}{
		embed:     embed(*w),
		UpdatedAt: core.NewDateTime(w.UpdatedAt),
		CreatedAt: core.NewDateTime(w.CreatedAt),
		ExpiredAt: core.NewOptionalDateTime(w.ExpiredAt),
	}
	return json.Marshal(marshaler)
}

func (w *Workbook) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

// Settings for a workbook
type WorkbookConfigSettings struct {
	// Whether to track changes for this workbook. Defaults to false. Tracking changes on a workbook allows for disabling workbook and sheet actions while data in the workbook is still being processed. You must run a recordHook listener if you enable this feature.
	TrackChanges *bool `json:"trackChanges,omitempty" url:"trackChanges,omitempty"`
	// When noMappingRedirect is set to true, dragging a file into a sheet will not redirect to the mapping screen. Defaults to false.
	NoMappingRedirect *bool `json:"noMappingRedirect,omitempty" url:"noMappingRedirect,omitempty"`
	// Used to set the order of sheets in the sidebar. Sheets that are not specified will be shown after those listed.
	SheetSidebarOrder []SheetId `json:"sheetSidebarOrder,omitempty" url:"sheetSidebarOrder,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (w *WorkbookConfigSettings) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WorkbookConfigSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkbookConfigSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkbookConfigSettings(value)

	extraProperties, err := core.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties

	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WorkbookConfigSettings) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WorkbookResponse struct {
	Data *Workbook `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (w *WorkbookResponse) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WorkbookResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkbookResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkbookResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties

	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WorkbookResponse) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

// Available treatments for a workbook
type WorkbookTreatments string

const (
	WorkbookTreatmentsExtractedFromSource WorkbookTreatments = "EXTRACTED_FROM_SOURCE"
	WorkbookTreatmentsSmallData           WorkbookTreatments = "SMALL_DATA"
)

func NewWorkbookTreatmentsFromString(s string) (WorkbookTreatments, error) {
	switch s {
	case "EXTRACTED_FROM_SOURCE":
		return WorkbookTreatmentsExtractedFromSource, nil
	case "SMALL_DATA":
		return WorkbookTreatmentsSmallData, nil
	}
	var t WorkbookTreatments
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (w WorkbookTreatments) Ptr() *WorkbookTreatments {
	return &w
}

// The updates to be made to an existing workbook
type WorkbookUpdate struct {
	// The name of the Workbook.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// An optional list of labels for the Workbook.
	Labels []string `json:"labels,omitempty" url:"labels,omitempty"`
	// The Space Id associated with the Workbook.
	SpaceId *SpaceId `json:"spaceId,omitempty" url:"spaceId,omitempty"`
	// The Environment Id associated with the Workbook.
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty" url:"environmentId,omitempty"`
	// The namespace of the Workbook.
	Namespace *string `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Describes shape of data as well as behavior
	Sheets  []*SheetConfigOrUpdate `json:"sheets,omitempty" url:"sheets,omitempty"`
	Actions []*Action              `json:"actions,omitempty" url:"actions,omitempty"`
	// Metadata for the workbook
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (w *WorkbookUpdate) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WorkbookUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkbookUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkbookUpdate(value)

	extraProperties, err := core.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties

	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WorkbookUpdate) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}
