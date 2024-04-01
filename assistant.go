// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type ListPromptsRequest struct {
	// Number of prompts to return in a page (default 7)
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// Based on pageSize, which page of prompts to return
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
}

// Create a prompts
type PromptCreate struct {
	Prompt string `json:"prompt" url:"prompt"`

	_rawJSON json.RawMessage
}

func (p *PromptCreate) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptCreate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptCreate(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptCreate) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// Update a prompts
type PromptPatch struct {
	Prompt *string `json:"prompt,omitempty" url:"prompt,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PromptPatch) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptPatch
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptPatch(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptPatch) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PromptResponse struct {
	Data *Prompt `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PromptResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptResponse(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PromptsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty" url:"pagination,omitempty"`
	Data       []*Prompt   `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PromptsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptsResponse(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptsResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}
