// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/FlatFilers/flatfile-go/internal"
	time "time"
)

type DeleteAllHistoryForUserRequest struct {
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty" url:"-"`
}

type DeleteMultipleRulesRequest struct {
	// Array of rule IDs to be deleted
	RuleIds []MappingId `json:"ruleIds,omitempty" url:"-"`
}

type ListProgramsRequest struct {
	// Number of programs to return in a page (default 10)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of records to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
	// Filter by creator
	CreatedBy *UserId `json:"-" url:"createdBy,omitempty"`
	// Filter by creation time
	CreatedAfter *time.Time `json:"-" url:"createdAfter,omitempty"`
	// Filter by creation time
	CreatedBefore *time.Time `json:"-" url:"createdBefore,omitempty"`
	// The ID of the environment
	EnvironmentId *EnvironmentId `json:"-" url:"environmentId,omitempty"`
	// Filter by family
	FamilyId *FamilyId `json:"-" url:"familyId,omitempty"`
	// Filter by namespace
	Namespace *string `json:"-" url:"namespace,omitempty"`
	// Filter by source keys
	SourceKeys *string `json:"-" url:"sourceKeys,omitempty"`
	// Filter by destination keys
	DestinationKeys *string `json:"-" url:"destinationKeys,omitempty"`
}

type CreateMappingRulesRequest = []*MappingRuleConfig

type MappingRule struct {
	// Name of the mapping rule
	Name   string      `json:"name" url:"name"`
	Type   string      `json:"type" url:"type"`
	Config interface{} `json:"config,omitempty" url:"config,omitempty"`
	// Time the mapping rule was last updated
	AcceptedAt *time.Time `json:"acceptedAt,omitempty" url:"acceptedAt,omitempty"`
	// User ID of the contributor of the mapping rule
	AcceptedBy *UserId `json:"acceptedBy,omitempty" url:"acceptedBy,omitempty"`
	// Metadata of the mapping rule
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// ID of the mapping rule
	Id MappingId `json:"id" url:"id"`
	// Confidence of the mapping rule
	Confidence *int `json:"confidence,omitempty" url:"confidence,omitempty"`
	// User ID of the user who suggested the mapping rule
	CreatedBy *UserId `json:"createdBy,omitempty" url:"createdBy,omitempty"`
	// Time the mapping rule was created
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// Time the mapping rule was last updated
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`
	// Time the mapping rule was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty" url:"deletedAt,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MappingRule) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

func (m *MappingRule) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}

func (m *MappingRule) GetConfig() interface{} {
	if m == nil {
		return nil
	}
	return m.Config
}

func (m *MappingRule) GetAcceptedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.AcceptedAt
}

func (m *MappingRule) GetAcceptedBy() *UserId {
	if m == nil {
		return nil
	}
	return m.AcceptedBy
}

func (m *MappingRule) GetMetadata() interface{} {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MappingRule) GetId() MappingId {
	if m == nil {
		return ""
	}
	return m.Id
}

func (m *MappingRule) GetConfidence() *int {
	if m == nil {
		return nil
	}
	return m.Confidence
}

func (m *MappingRule) GetCreatedBy() *UserId {
	if m == nil {
		return nil
	}
	return m.CreatedBy
}

func (m *MappingRule) GetCreatedAt() time.Time {
	if m == nil {
		return time.Time{}
	}
	return m.CreatedAt
}

func (m *MappingRule) GetUpdatedAt() time.Time {
	if m == nil {
		return time.Time{}
	}
	return m.UpdatedAt
}

func (m *MappingRule) GetDeletedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.DeletedAt
}

func (m *MappingRule) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MappingRule) UnmarshalJSON(data []byte) error {
	type embed MappingRule
	var unmarshaler = struct {
		embed
		AcceptedAt *internal.DateTime `json:"acceptedAt,omitempty"`
		CreatedAt  *internal.DateTime `json:"createdAt"`
		UpdatedAt  *internal.DateTime `json:"updatedAt"`
		DeletedAt  *internal.DateTime `json:"deletedAt,omitempty"`
	}{
		embed: embed(*m),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*m = MappingRule(unmarshaler.embed)
	m.AcceptedAt = unmarshaler.AcceptedAt.TimePtr()
	m.CreatedAt = unmarshaler.CreatedAt.Time()
	m.UpdatedAt = unmarshaler.UpdatedAt.Time()
	m.DeletedAt = unmarshaler.DeletedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MappingRule) MarshalJSON() ([]byte, error) {
	type embed MappingRule
	var marshaler = struct {
		embed
		AcceptedAt *internal.DateTime `json:"acceptedAt,omitempty"`
		CreatedAt  *internal.DateTime `json:"createdAt"`
		UpdatedAt  *internal.DateTime `json:"updatedAt"`
		DeletedAt  *internal.DateTime `json:"deletedAt,omitempty"`
	}{
		embed:      embed(*m),
		AcceptedAt: internal.NewOptionalDateTime(m.AcceptedAt),
		CreatedAt:  internal.NewDateTime(m.CreatedAt),
		UpdatedAt:  internal.NewDateTime(m.UpdatedAt),
		DeletedAt:  internal.NewOptionalDateTime(m.DeletedAt),
	}
	return json.Marshal(marshaler)
}

func (m *MappingRule) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MappingRuleConfig struct {
	// Name of the mapping rule
	Name   string      `json:"name" url:"name"`
	Type   string      `json:"type" url:"type"`
	Config interface{} `json:"config,omitempty" url:"config,omitempty"`
	// Time the mapping rule was last updated
	AcceptedAt *time.Time `json:"acceptedAt,omitempty" url:"acceptedAt,omitempty"`
	// User ID of the contributor of the mapping rule
	AcceptedBy *UserId `json:"acceptedBy,omitempty" url:"acceptedBy,omitempty"`
	// Metadata of the mapping rule
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MappingRuleConfig) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

func (m *MappingRuleConfig) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}

func (m *MappingRuleConfig) GetConfig() interface{} {
	if m == nil {
		return nil
	}
	return m.Config
}

func (m *MappingRuleConfig) GetAcceptedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.AcceptedAt
}

func (m *MappingRuleConfig) GetAcceptedBy() *UserId {
	if m == nil {
		return nil
	}
	return m.AcceptedBy
}

func (m *MappingRuleConfig) GetMetadata() interface{} {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MappingRuleConfig) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MappingRuleConfig) UnmarshalJSON(data []byte) error {
	type embed MappingRuleConfig
	var unmarshaler = struct {
		embed
		AcceptedAt *internal.DateTime `json:"acceptedAt,omitempty"`
	}{
		embed: embed(*m),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*m = MappingRuleConfig(unmarshaler.embed)
	m.AcceptedAt = unmarshaler.AcceptedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MappingRuleConfig) MarshalJSON() ([]byte, error) {
	type embed MappingRuleConfig
	var marshaler = struct {
		embed
		AcceptedAt *internal.DateTime `json:"acceptedAt,omitempty"`
	}{
		embed:      embed(*m),
		AcceptedAt: internal.NewOptionalDateTime(m.AcceptedAt),
	}
	return json.Marshal(marshaler)
}

func (m *MappingRuleConfig) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MappingRuleOrConfig struct {
	// Name of the mapping rule
	Name   string      `json:"name" url:"name"`
	Type   string      `json:"type" url:"type"`
	Config interface{} `json:"config,omitempty" url:"config,omitempty"`
	// Time the mapping rule was last updated
	AcceptedAt *time.Time `json:"acceptedAt,omitempty" url:"acceptedAt,omitempty"`
	// User ID of the contributor of the mapping rule
	AcceptedBy *UserId `json:"acceptedBy,omitempty" url:"acceptedBy,omitempty"`
	// Metadata of the mapping rule
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// ID of the mapping rule
	Id *MappingId `json:"id,omitempty" url:"id,omitempty"`
	// Confidence of the mapping rule
	Confidence *int `json:"confidence,omitempty" url:"confidence,omitempty"`
	// User ID of the creator of the mapping rule
	CreatedBy *UserId `json:"createdBy,omitempty" url:"createdBy,omitempty"`
	// Time the mapping rule was created
	CreatedAt *time.Time `json:"createdAt,omitempty" url:"createdAt,omitempty"`
	// Time the mapping rule was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty" url:"updatedAt,omitempty"`
	// Time the mapping rule was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty" url:"deletedAt,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MappingRuleOrConfig) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

func (m *MappingRuleOrConfig) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}

func (m *MappingRuleOrConfig) GetConfig() interface{} {
	if m == nil {
		return nil
	}
	return m.Config
}

func (m *MappingRuleOrConfig) GetAcceptedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.AcceptedAt
}

func (m *MappingRuleOrConfig) GetAcceptedBy() *UserId {
	if m == nil {
		return nil
	}
	return m.AcceptedBy
}

func (m *MappingRuleOrConfig) GetMetadata() interface{} {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MappingRuleOrConfig) GetId() *MappingId {
	if m == nil {
		return nil
	}
	return m.Id
}

func (m *MappingRuleOrConfig) GetConfidence() *int {
	if m == nil {
		return nil
	}
	return m.Confidence
}

func (m *MappingRuleOrConfig) GetCreatedBy() *UserId {
	if m == nil {
		return nil
	}
	return m.CreatedBy
}

func (m *MappingRuleOrConfig) GetCreatedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.CreatedAt
}

func (m *MappingRuleOrConfig) GetUpdatedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.UpdatedAt
}

func (m *MappingRuleOrConfig) GetDeletedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.DeletedAt
}

func (m *MappingRuleOrConfig) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MappingRuleOrConfig) UnmarshalJSON(data []byte) error {
	type embed MappingRuleOrConfig
	var unmarshaler = struct {
		embed
		AcceptedAt *internal.DateTime `json:"acceptedAt,omitempty"`
		CreatedAt  *internal.DateTime `json:"createdAt,omitempty"`
		UpdatedAt  *internal.DateTime `json:"updatedAt,omitempty"`
		DeletedAt  *internal.DateTime `json:"deletedAt,omitempty"`
	}{
		embed: embed(*m),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*m = MappingRuleOrConfig(unmarshaler.embed)
	m.AcceptedAt = unmarshaler.AcceptedAt.TimePtr()
	m.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	m.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	m.DeletedAt = unmarshaler.DeletedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MappingRuleOrConfig) MarshalJSON() ([]byte, error) {
	type embed MappingRuleOrConfig
	var marshaler = struct {
		embed
		AcceptedAt *internal.DateTime `json:"acceptedAt,omitempty"`
		CreatedAt  *internal.DateTime `json:"createdAt,omitempty"`
		UpdatedAt  *internal.DateTime `json:"updatedAt,omitempty"`
		DeletedAt  *internal.DateTime `json:"deletedAt,omitempty"`
	}{
		embed:      embed(*m),
		AcceptedAt: internal.NewOptionalDateTime(m.AcceptedAt),
		CreatedAt:  internal.NewOptionalDateTime(m.CreatedAt),
		UpdatedAt:  internal.NewOptionalDateTime(m.UpdatedAt),
		DeletedAt:  internal.NewOptionalDateTime(m.DeletedAt),
	}
	return json.Marshal(marshaler)
}

func (m *MappingRuleOrConfig) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MappingRuleResponse struct {
	Data *MappingRule `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MappingRuleResponse) GetData() *MappingRule {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MappingRuleResponse) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MappingRuleResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler MappingRuleResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MappingRuleResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MappingRuleResponse) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MappingRulesResponse struct {
	Data []*MappingRule `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MappingRulesResponse) GetData() []*MappingRule {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MappingRulesResponse) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MappingRulesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler MappingRulesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MappingRulesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MappingRulesResponse) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type Program struct {
	// Mapping rules
	Rules []*MappingRuleOrConfig `json:"rules,omitempty" url:"rules,omitempty"`
	// If this program was saved, this is the ID of the program
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// Namespace of the program
	Namespace *string `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Family ID of the program, if it belongs to a family
	FamilyId *FamilyId `json:"familyId,omitempty" url:"familyId,omitempty"`
	// If this program was saved, this is the time it was created
	CreatedAt *time.Time `json:"createdAt,omitempty" url:"createdAt,omitempty"`
	// If this program was saved, this is the user ID of the creator
	CreatedBy *UserId `json:"createdBy,omitempty" url:"createdBy,omitempty"`
	// Source keys
	SourceKeys []string `json:"sourceKeys,omitempty" url:"sourceKeys,omitempty"`
	// Destination keys
	DestinationKeys []string `json:"destinationKeys,omitempty" url:"destinationKeys,omitempty"`
	// Summary of the mapping rules
	Summary *ProgramSummary `json:"summary,omitempty" url:"summary,omitempty"`
	// If this program was saved, this token allows you to modify the program
	AccessToken *string `json:"accessToken,omitempty" url:"accessToken,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *Program) GetRules() []*MappingRuleOrConfig {
	if p == nil {
		return nil
	}
	return p.Rules
}

func (p *Program) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Program) GetNamespace() *string {
	if p == nil {
		return nil
	}
	return p.Namespace
}

func (p *Program) GetFamilyId() *FamilyId {
	if p == nil {
		return nil
	}
	return p.FamilyId
}

func (p *Program) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *Program) GetCreatedBy() *UserId {
	if p == nil {
		return nil
	}
	return p.CreatedBy
}

func (p *Program) GetSourceKeys() []string {
	if p == nil {
		return nil
	}
	return p.SourceKeys
}

func (p *Program) GetDestinationKeys() []string {
	if p == nil {
		return nil
	}
	return p.DestinationKeys
}

func (p *Program) GetSummary() *ProgramSummary {
	if p == nil {
		return nil
	}
	return p.Summary
}

func (p *Program) GetAccessToken() *string {
	if p == nil {
		return nil
	}
	return p.AccessToken
}

func (p *Program) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *Program) UnmarshalJSON(data []byte) error {
	type embed Program
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = Program(unmarshaler.embed)
	p.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *Program) MarshalJSON() ([]byte, error) {
	type embed Program
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt,omitempty"`
	}{
		embed:     embed(*p),
		CreatedAt: internal.NewOptionalDateTime(p.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (p *Program) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type ProgramConfig struct {
	// Source schema
	Source *SheetConfig `json:"source,omitempty" url:"source,omitempty"`
	// Destination schema
	Destination *SheetConfig `json:"destination,omitempty" url:"destination,omitempty"`
	// ID of the family to add the program to
	FamilyId *FamilyId `json:"familyId,omitempty" url:"familyId,omitempty"`
	// Namespace of the program
	Namespace *string `json:"namespace,omitempty" url:"namespace,omitempty"`
	// Whether to save the program for editing later. Defaults to false. If true, the response will contain an ID and access token.
	Save *bool `json:"save,omitempty" url:"save,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *ProgramConfig) GetSource() *SheetConfig {
	if p == nil {
		return nil
	}
	return p.Source
}

func (p *ProgramConfig) GetDestination() *SheetConfig {
	if p == nil {
		return nil
	}
	return p.Destination
}

func (p *ProgramConfig) GetFamilyId() *FamilyId {
	if p == nil {
		return nil
	}
	return p.FamilyId
}

func (p *ProgramConfig) GetNamespace() *string {
	if p == nil {
		return nil
	}
	return p.Namespace
}

func (p *ProgramConfig) GetSave() *bool {
	if p == nil {
		return nil
	}
	return p.Save
}

func (p *ProgramConfig) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProgramConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ProgramConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProgramConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProgramConfig) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type ProgramResponse struct {
	Data *Program `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *ProgramResponse) GetData() *Program {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *ProgramResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProgramResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ProgramResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProgramResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProgramResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type ProgramSummary struct {
	// Total number of mapping rules
	TotalRuleCount int `json:"totalRuleCount" url:"totalRuleCount"`
	// Number of mapping rules added
	AddedRuleCount int `json:"addedRuleCount" url:"addedRuleCount"`
	// Number of mapping rules deleted
	DeletedRuleCount int `json:"deletedRuleCount" url:"deletedRuleCount"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *ProgramSummary) GetTotalRuleCount() int {
	if p == nil {
		return 0
	}
	return p.TotalRuleCount
}

func (p *ProgramSummary) GetAddedRuleCount() int {
	if p == nil {
		return 0
	}
	return p.AddedRuleCount
}

func (p *ProgramSummary) GetDeletedRuleCount() int {
	if p == nil {
		return 0
	}
	return p.DeletedRuleCount
}

func (p *ProgramSummary) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProgramSummary) UnmarshalJSON(data []byte) error {
	type unmarshaler ProgramSummary
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProgramSummary(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProgramSummary) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type ProgramsResponse struct {
	Data []*Program `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *ProgramsResponse) GetData() []*Program {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *ProgramsResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProgramsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ProgramsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProgramsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProgramsResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type UpdateMappingRulesRequest = []*MappingRule
