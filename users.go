// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/FlatFilers/flatfile-go/internal"
	time "time"
)

type ListUsersRequest struct {
	// Email of guest to return
	Email *string `json:"-" url:"email,omitempty"`
	// String to search for users by name and email
	Search *string `json:"-" url:"search,omitempty"`
	// Field to sort users by
	SortField *ListUsersSortField `json:"-" url:"sortField,omitempty"`
	// Direction of sorting
	SortDirection *SortDirection `json:"-" url:"sortDirection,omitempty"`
	// Number of users to return in a page (default 20)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of users to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
}

type ListUsersResponse struct {
	Data       []*User     `json:"data,omitempty" url:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty" url:"pagination,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListUsersResponse) GetData() []*User {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListUsersResponse) GetPagination() *Pagination {
	if l == nil {
		return nil
	}
	return l.Pagination
}

func (l *ListUsersResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListUsersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListUsersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListUsersResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListUsersResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListUsersSortField string

const (
	ListUsersSortFieldEmail      ListUsersSortField = "email"
	ListUsersSortFieldName       ListUsersSortField = "name"
	ListUsersSortFieldId         ListUsersSortField = "id"
	ListUsersSortFieldIdp        ListUsersSortField = "idp"
	ListUsersSortFieldIdpRef     ListUsersSortField = "idp_ref"
	ListUsersSortFieldCreatedAt  ListUsersSortField = "created_at"
	ListUsersSortFieldUpdatedAt  ListUsersSortField = "updated_at"
	ListUsersSortFieldLastSeenAt ListUsersSortField = "last_seen_at"
	ListUsersSortFieldDashboard  ListUsersSortField = "dashboard"
)

func NewListUsersSortFieldFromString(s string) (ListUsersSortField, error) {
	switch s {
	case "email":
		return ListUsersSortFieldEmail, nil
	case "name":
		return ListUsersSortFieldName, nil
	case "id":
		return ListUsersSortFieldId, nil
	case "idp":
		return ListUsersSortFieldIdp, nil
	case "idp_ref":
		return ListUsersSortFieldIdpRef, nil
	case "created_at":
		return ListUsersSortFieldCreatedAt, nil
	case "updated_at":
		return ListUsersSortFieldUpdatedAt, nil
	case "last_seen_at":
		return ListUsersSortFieldLastSeenAt, nil
	case "dashboard":
		return ListUsersSortFieldDashboard, nil
	}
	var t ListUsersSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListUsersSortField) Ptr() *ListUsersSortField {
	return &l
}

// Configurations for the user
type User struct {
	Email      string                 `json:"email" url:"email"`
	Name       string                 `json:"name" url:"name"`
	AccountId  AccountId              `json:"accountId" url:"accountId"`
	Id         UserId                 `json:"id" url:"id"`
	Idp        string                 `json:"idp" url:"idp"`
	IdpRef     *string                `json:"idpRef,omitempty" url:"idpRef,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	CreatedAt  time.Time              `json:"createdAt" url:"createdAt"`
	UpdatedAt  time.Time              `json:"updatedAt" url:"updatedAt"`
	LastSeenAt *time.Time             `json:"lastSeenAt,omitempty" url:"lastSeenAt,omitempty"`
	Dashboard  *int                   `json:"dashboard,omitempty" url:"dashboard,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *User) GetEmail() string {
	if u == nil {
		return ""
	}
	return u.Email
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *User) GetAccountId() AccountId {
	if u == nil {
		return ""
	}
	return u.AccountId
}

func (u *User) GetId() UserId {
	if u == nil {
		return ""
	}
	return u.Id
}

func (u *User) GetIdp() string {
	if u == nil {
		return ""
	}
	return u.Idp
}

func (u *User) GetIdpRef() *string {
	if u == nil {
		return nil
	}
	return u.IdpRef
}

func (u *User) GetMetadata() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.Metadata
}

func (u *User) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *User) GetLastSeenAt() *time.Time {
	if u == nil {
		return nil
	}
	return u.LastSeenAt
}

func (u *User) GetDashboard() *int {
	if u == nil {
		return nil
	}
	return u.Dashboard
}

func (u *User) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *User) UnmarshalJSON(data []byte) error {
	type embed User
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"createdAt"`
		UpdatedAt  *internal.DateTime `json:"updatedAt"`
		LastSeenAt *internal.DateTime `json:"lastSeenAt,omitempty"`
	}{
		embed: embed(*u),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*u = User(unmarshaler.embed)
	u.CreatedAt = unmarshaler.CreatedAt.Time()
	u.UpdatedAt = unmarshaler.UpdatedAt.Time()
	u.LastSeenAt = unmarshaler.LastSeenAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	type embed User
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"createdAt"`
		UpdatedAt  *internal.DateTime `json:"updatedAt"`
		LastSeenAt *internal.DateTime `json:"lastSeenAt,omitempty"`
	}{
		embed:      embed(*u),
		CreatedAt:  internal.NewDateTime(u.CreatedAt),
		UpdatedAt:  internal.NewDateTime(u.UpdatedAt),
		LastSeenAt: internal.NewOptionalDateTime(u.LastSeenAt),
	}
	return json.Marshal(marshaler)
}

func (u *User) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

// Properties used to create a new user
type UserConfig struct {
	Email     string    `json:"email" url:"email"`
	Name      string    `json:"name" url:"name"`
	AccountId AccountId `json:"accountId" url:"accountId"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UserConfig) GetEmail() string {
	if u == nil {
		return ""
	}
	return u.Email
}

func (u *UserConfig) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *UserConfig) GetAccountId() AccountId {
	if u == nil {
		return ""
	}
	return u.AccountId
}

func (u *UserConfig) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler UserConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserConfig) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

// Properties used to create a new user
type UserCreateAndInviteRequest struct {
	Email      string                    `json:"email" url:"email"`
	Name       string                    `json:"name" url:"name"`
	ActorRoles []*AssignActorRoleRequest `json:"actorRoles,omitempty" url:"actorRoles,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UserCreateAndInviteRequest) GetEmail() string {
	if u == nil {
		return ""
	}
	return u.Email
}

func (u *UserCreateAndInviteRequest) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *UserCreateAndInviteRequest) GetActorRoles() []*AssignActorRoleRequest {
	if u == nil {
		return nil
	}
	return u.ActorRoles
}

func (u *UserCreateAndInviteRequest) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserCreateAndInviteRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UserCreateAndInviteRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserCreateAndInviteRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserCreateAndInviteRequest) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserResponse struct {
	Data *User `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UserResponse) GetData() *User {
	if u == nil {
		return nil
	}
	return u.Data
}

func (u *UserResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserResponse) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserWithRoles struct {
	Email      string                 `json:"email" url:"email"`
	Name       string                 `json:"name" url:"name"`
	AccountId  AccountId              `json:"accountId" url:"accountId"`
	Id         UserId                 `json:"id" url:"id"`
	Idp        string                 `json:"idp" url:"idp"`
	IdpRef     *string                `json:"idpRef,omitempty" url:"idpRef,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	CreatedAt  time.Time              `json:"createdAt" url:"createdAt"`
	UpdatedAt  time.Time              `json:"updatedAt" url:"updatedAt"`
	LastSeenAt *time.Time             `json:"lastSeenAt,omitempty" url:"lastSeenAt,omitempty"`
	Dashboard  *int                   `json:"dashboard,omitempty" url:"dashboard,omitempty"`
	ActorRoles []*ActorRoleResponse   `json:"actorRoles,omitempty" url:"actorRoles,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UserWithRoles) GetEmail() string {
	if u == nil {
		return ""
	}
	return u.Email
}

func (u *UserWithRoles) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *UserWithRoles) GetAccountId() AccountId {
	if u == nil {
		return ""
	}
	return u.AccountId
}

func (u *UserWithRoles) GetId() UserId {
	if u == nil {
		return ""
	}
	return u.Id
}

func (u *UserWithRoles) GetIdp() string {
	if u == nil {
		return ""
	}
	return u.Idp
}

func (u *UserWithRoles) GetIdpRef() *string {
	if u == nil {
		return nil
	}
	return u.IdpRef
}

func (u *UserWithRoles) GetMetadata() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.Metadata
}

func (u *UserWithRoles) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *UserWithRoles) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *UserWithRoles) GetLastSeenAt() *time.Time {
	if u == nil {
		return nil
	}
	return u.LastSeenAt
}

func (u *UserWithRoles) GetDashboard() *int {
	if u == nil {
		return nil
	}
	return u.Dashboard
}

func (u *UserWithRoles) GetActorRoles() []*ActorRoleResponse {
	if u == nil {
		return nil
	}
	return u.ActorRoles
}

func (u *UserWithRoles) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserWithRoles) UnmarshalJSON(data []byte) error {
	type embed UserWithRoles
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"createdAt"`
		UpdatedAt  *internal.DateTime `json:"updatedAt"`
		LastSeenAt *internal.DateTime `json:"lastSeenAt,omitempty"`
	}{
		embed: embed(*u),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*u = UserWithRoles(unmarshaler.embed)
	u.CreatedAt = unmarshaler.CreatedAt.Time()
	u.UpdatedAt = unmarshaler.UpdatedAt.Time()
	u.LastSeenAt = unmarshaler.LastSeenAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserWithRoles) MarshalJSON() ([]byte, error) {
	type embed UserWithRoles
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"createdAt"`
		UpdatedAt  *internal.DateTime `json:"updatedAt"`
		LastSeenAt *internal.DateTime `json:"lastSeenAt,omitempty"`
	}{
		embed:      embed(*u),
		CreatedAt:  internal.NewDateTime(u.CreatedAt),
		UpdatedAt:  internal.NewDateTime(u.UpdatedAt),
		LastSeenAt: internal.NewOptionalDateTime(u.LastSeenAt),
	}
	return json.Marshal(marshaler)
}

func (u *UserWithRoles) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserWithRolesResponse struct {
	Data *UserWithRoles `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UserWithRolesResponse) GetData() *UserWithRoles {
	if u == nil {
		return nil
	}
	return u.Data
}

func (u *UserWithRolesResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserWithRolesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserWithRolesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserWithRolesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserWithRolesResponse) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdateUserRequest struct {
	Name      *string `json:"name,omitempty" url:"-"`
	Dashboard *int    `json:"dashboard,omitempty" url:"-"`
}
