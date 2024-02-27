// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type ListJobsRequest struct {
	// When provided, only jobs for the given environment will be returned
	EnvironmentId *EnvironmentId `json:"-" url:"environmentId,omitempty"`
	// When provided, only jobs for the given space will be returned
	SpaceId *SpaceId `json:"-" url:"spaceId,omitempty"`
	// When provided, only jobs for the given workbook will be returned
	WorkbookId *WorkbookId `json:"-" url:"workbookId,omitempty"`
	// When provided, only jobs for the given file will be returned
	FileId *FileId `json:"-" url:"fileId,omitempty"`
	// When provided, only jobs that are parts of the given job will be returned
	ParentId *JobId `json:"-" url:"parentId,omitempty"`
	// Number of jobs to return in a page (default 20)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of jobs to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
	// Sort direction - asc (ascending) or desc (descending)
	SortDirection *SortDirection `json:"-" url:"sortDirection,omitempty"`
}

// Pipeline Job ID
type JobId = string

// Details about the user who acknowledged the job
type JobAckDetails struct {
	Info *string `json:"info,omitempty" url:"info,omitempty"`
	// the progress of the job. Whole number between 0 and 100
	Progress              *int       `json:"progress,omitempty" url:"progress,omitempty"`
	EstimatedCompletionAt *time.Time `json:"estimatedCompletionAt,omitempty" url:"estimatedCompletionAt,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobAckDetails) UnmarshalJSON(data []byte) error {
	type embed JobAckDetails
	var unmarshaler = struct {
		embed
		EstimatedCompletionAt *core.DateTime `json:"estimatedCompletionAt,omitempty"`
	}{
		embed: embed(*j),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*j = JobAckDetails(unmarshaler.embed)
	j.EstimatedCompletionAt = unmarshaler.EstimatedCompletionAt.TimePtr()
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobAckDetails) MarshalJSON() ([]byte, error) {
	type embed JobAckDetails
	var marshaler = struct {
		embed
		EstimatedCompletionAt *core.DateTime `json:"estimatedCompletionAt,omitempty"`
	}{
		embed:                 embed(*j),
		EstimatedCompletionAt: core.NewOptionalDateTime(j.EstimatedCompletionAt),
	}
	return json.Marshal(marshaler)
}

func (j *JobAckDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// Info about the reason the job was canceled
type JobCancelDetails struct {
	Info *string `json:"info,omitempty" url:"info,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobCancelDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobCancelDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobCancelDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobCancelDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// Outcome summary of a job
type JobCompleteDetails struct {
	Outcome *JobOutcome `json:"outcome,omitempty" url:"outcome,omitempty"`
	Info    *string     `json:"info,omitempty" url:"info,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobCompleteDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobCompleteDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobCompleteDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobCompleteDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// A single unit of work that a pipeline will execute
type JobConfig struct {
	// The type of job
	Type JobType `json:"type,omitempty" url:"type,omitempty"`
	// the type of operation to perform on the data. For example, "export".
	Operation   string           `json:"operation" url:"operation"`
	Source      JobSource        `json:"source" url:"source"`
	Destination *JobDestination  `json:"destination,omitempty" url:"destination,omitempty"`
	Config      *JobUpdateConfig `json:"config,omitempty" url:"config,omitempty"`
	// the type of trigger to use for this job
	Trigger *Trigger `json:"trigger,omitempty" url:"trigger,omitempty"`
	// the status of the job
	Status *JobStatus `json:"status,omitempty" url:"status,omitempty"`
	// the progress of the job. Whole number between 0 and 100
	Progress *int    `json:"progress,omitempty" url:"progress,omitempty"`
	FileId   *FileId `json:"fileId,omitempty" url:"fileId,omitempty"`
	// the mode of the job
	Mode *JobMode `json:"mode,omitempty" url:"mode,omitempty"`
	// Input parameters for this job type.
	Input map[string]interface{} `json:"input,omitempty" url:"input,omitempty"`
	// Subject parameters for this job type.
	Subject *JobSubject `json:"subject,omitempty" url:"subject,omitempty"`
	// Outcome summary of job.
	Outcome map[string]interface{} `json:"outcome,omitempty" url:"outcome,omitempty"`
	// Current status of job in text
	Info *string `json:"info,omitempty" url:"info,omitempty"`
	// Indicates if Flatfile is managing the control flow of this job or if it is being manually tracked.
	Managed *bool `json:"managed,omitempty" url:"managed,omitempty"`
	// The id of the environment this job belongs to
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty" url:"environmentId,omitempty"`
	// The part number of this job
	Part *int `json:"part,omitempty" url:"part,omitempty"`
	// The data for this part of the job
	PartData map[string]interface{} `json:"partData,omitempty" url:"partData,omitempty"`
	// The execution mode for this part of the job
	PartExecution *JobPartExecution `json:"partExecution,omitempty" url:"partExecution,omitempty"`
	// The id of the parent job
	ParentId *JobId `json:"parentId,omitempty" url:"parentId,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler JobConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobConfig(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobConfig) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobExecutionPlanConfigRequest struct {
	FieldMapping              []*Edge             `json:"fieldMapping,omitempty" url:"fieldMapping,omitempty"`
	UnmappedSourceFields      []*SourceField      `json:"unmappedSourceFields,omitempty" url:"unmappedSourceFields,omitempty"`
	UnmappedDestinationFields []*DestinationField `json:"unmappedDestinationFields,omitempty" url:"unmappedDestinationFields,omitempty"`
	ProgramId                 *string             `json:"programId,omitempty" url:"programId,omitempty"`
	FileId                    FileId              `json:"fileId" url:"fileId"`
	JobId                     JobId               `json:"jobId" url:"jobId"`

	_rawJSON json.RawMessage
}

func (j *JobExecutionPlanConfigRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler JobExecutionPlanConfigRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobExecutionPlanConfigRequest(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobExecutionPlanConfigRequest) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobExecutionPlanRequest struct {
	FieldMapping              []*Edge             `json:"fieldMapping,omitempty" url:"fieldMapping,omitempty"`
	UnmappedSourceFields      []*SourceField      `json:"unmappedSourceFields,omitempty" url:"unmappedSourceFields,omitempty"`
	UnmappedDestinationFields []*DestinationField `json:"unmappedDestinationFields,omitempty" url:"unmappedDestinationFields,omitempty"`
	ProgramId                 *string             `json:"programId,omitempty" url:"programId,omitempty"`
	FileId                    FileId              `json:"fileId" url:"fileId"`
	JobId                     JobId               `json:"jobId" url:"jobId"`

	_rawJSON json.RawMessage
}

func (j *JobExecutionPlanRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler JobExecutionPlanRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobExecutionPlanRequest(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobExecutionPlanRequest) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobPlanResponse struct {
	Data *JobPlan `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobPlanResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler JobPlanResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobPlanResponse(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobPlanResponse) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobResponse struct {
	Data *Job `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler JobResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobResponse(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobResponse) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// Info about the reason the job was split
type JobSplitDetails struct {
	Parts         *JobParts `json:"parts,omitempty" url:"parts,omitempty"`
	RunInParallel *bool     `json:"runInParallel,omitempty" url:"runInParallel,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobSplitDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobSplitDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobSplitDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobSplitDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// A single unit of work that will be executed
type JobUpdate struct {
	Config *JobUpdateConfig `json:"config,omitempty" url:"config,omitempty"`
	// the status of the job
	Status *JobStatus `json:"status,omitempty" url:"status,omitempty"`
	// the progress of the job. Whole number between 0 and 100
	Progress *int `json:"progress,omitempty" url:"progress,omitempty"`
	// the time that the job's outcome has been acknowledged by a user
	OutcomeAcknowledgedAt *time.Time `json:"outcomeAcknowledgedAt,omitempty" url:"outcomeAcknowledgedAt,omitempty"`
	// Current status of job in text
	Info *string `json:"info,omitempty" url:"info,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobUpdate) UnmarshalJSON(data []byte) error {
	type embed JobUpdate
	var unmarshaler = struct {
		embed
		OutcomeAcknowledgedAt *core.DateTime `json:"outcomeAcknowledgedAt,omitempty"`
	}{
		embed: embed(*j),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*j = JobUpdate(unmarshaler.embed)
	j.OutcomeAcknowledgedAt = unmarshaler.OutcomeAcknowledgedAt.TimePtr()
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobUpdate) MarshalJSON() ([]byte, error) {
	type embed JobUpdate
	var marshaler = struct {
		embed
		OutcomeAcknowledgedAt *core.DateTime `json:"outcomeAcknowledgedAt,omitempty"`
	}{
		embed:                 embed(*j),
		OutcomeAcknowledgedAt: core.NewOptionalDateTime(j.OutcomeAcknowledgedAt),
	}
	return json.Marshal(marshaler)
}

func (j *JobUpdate) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type ListJobsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty" url:"pagination,omitempty"`
	Data       []*Job      `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListJobsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListJobsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListJobsResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListJobsResponse) String() string {
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

type MutateJobConfig struct {
	SheetId SheetId `json:"sheetId" url:"sheetId"`
	// A JavaScript function that will be run on each record in the sheet, it should return a mutated record.
	MutateRecord string `json:"mutateRecord" url:"mutateRecord"`
	// If the mutation was generated through some sort of id-ed process, this links this job and that process.
	MutationId  *string      `json:"mutationId,omitempty" url:"mutationId,omitempty"`
	Filter      *Filter      `json:"filter,omitempty" url:"filter,omitempty"`
	FilterField *FilterField `json:"filterField,omitempty" url:"filterField,omitempty"`
	SearchValue *SearchValue `json:"searchValue,omitempty" url:"searchValue,omitempty"`
	SearchField *SearchField `json:"searchField,omitempty" url:"searchField,omitempty"`
	Q           *string      `json:"q,omitempty" url:"q,omitempty"`
	// The Record Ids param (ids) is a list of record ids that can be passed to several record endpoints allowing the user to identify specific records to INCLUDE in the query, or specific records to EXCLUDE, depending on whether or not filters are being applied. When passing a query param that filters the record dataset, such as 'searchValue', or a 'filter' of 'valid' | 'error' | 'all', the 'ids' param will EXCLUDE those records from the filtered results. For basic queries that do not filter the dataset, passing record ids in the 'ids' param will limit the dataset to INCLUDE just those specific records
	Ids []RecordId `json:"ids,omitempty" url:"ids,omitempty"`

	_rawJSON json.RawMessage
}

func (m *MutateJobConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler MutateJobConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MutateJobConfig(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MutateJobConfig) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}
