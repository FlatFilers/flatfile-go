// This file was auto-generated by Fern from our API Definition.

package agents

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	flatfilego "github.com/FlatFilers/flatfile-go"
	core "github.com/FlatFilers/flatfile-go/core"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

func (c *Client) List(ctx context.Context, request *flatfilego.ListAgentsRequest) (*flatfilego.ListAgentsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "agents"

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.ListAgentsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Create(ctx context.Context, request *flatfilego.CreateAgentsRequest) (*flatfilego.AgentResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "agents"

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.AgentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Get(ctx context.Context, agentId flatfilego.AgentId, request *flatfilego.GetAgentRequest) (*flatfilego.AgentResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"agents/%v", agentId)

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(flatfilego.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.AgentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetAgentLogs(ctx context.Context, agentId flatfilego.AgentId, request *flatfilego.GetAgentLogsRequest) (*flatfilego.GetAgentLogsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"agents/%v/logs", agentId)

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(flatfilego.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.GetAgentLogsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetAgentLog(ctx context.Context, eventId flatfilego.EventId, request *flatfilego.GetAgentLogRequest) (*flatfilego.GetDetailedAgentLogResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"agents/log/%v", eventId)

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(flatfilego.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.GetDetailedAgentLogResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetEnvironmentAgentLogs(ctx context.Context, request *flatfilego.GetEnvironmentAgentLogsRequest) (*flatfilego.GetDetailedAgentLogsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "agents/logs"

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if request.SpaceId != nil {
		queryParams.Add("spaceId", fmt.Sprintf("%v", request.SpaceId))
	}
	if request.Success != nil {
		queryParams.Add("success", fmt.Sprintf("%v", request.Success))
	}
	if request.PageSize != nil {
		queryParams.Add("pageSize", fmt.Sprintf("%v", request.PageSize))
	}
	if request.PageNumber != nil {
		queryParams.Add("pageNumber", fmt.Sprintf("%v", request.PageNumber))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(flatfilego.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.GetDetailedAgentLogsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetEnvironmentAgentExecutions(ctx context.Context, request *flatfilego.GetEnvironmentAgentExecutionsRequest) (*flatfilego.GetExecutionsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "agents/executions"

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if request.SpaceId != nil {
		queryParams.Add("spaceId", fmt.Sprintf("%v", request.SpaceId))
	}
	if request.Success != nil {
		queryParams.Add("success", fmt.Sprintf("%v", request.Success))
	}
	if request.PageSize != nil {
		queryParams.Add("pageSize", fmt.Sprintf("%v", request.PageSize))
	}
	if request.PageNumber != nil {
		queryParams.Add("pageNumber", fmt.Sprintf("%v", request.PageNumber))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(flatfilego.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.GetExecutionsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Deletes a single agent
func (c *Client) Delete(ctx context.Context, agentId flatfilego.AgentId, request *flatfilego.DeleteAgentRequest) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"agents/%v", agentId)

	queryParams := make(url.Values)
	queryParams.Add("environmentId", fmt.Sprintf("%v", request.EnvironmentId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(flatfilego.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(flatfilego.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodDelete,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
