// This file was auto-generated by Fern from our API Definition.

package records

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

// Returns records from a sheet in a workbook
//
// ID of sheet
func (c *Client) Get(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.GetRecordsRequest) (*flatfilego.GetRecordsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/records", sheetId)

	queryParams := make(url.Values)
	if request.VersionId != nil {
		queryParams.Add("versionId", fmt.Sprintf("%v", request.VersionId))
	}
	if request.CommitId != nil {
		queryParams.Add("commitId", fmt.Sprintf("%v", request.CommitId))
	}
	if request.SinceVersionId != nil {
		queryParams.Add("sinceVersionId", fmt.Sprintf("%v", request.SinceVersionId))
	}
	if request.SinceCommitId != nil {
		queryParams.Add("sinceCommitId", fmt.Sprintf("%v", request.SinceCommitId))
	}
	if request.SortField != nil {
		queryParams.Add("sortField", fmt.Sprintf("%v", request.SortField))
	}
	if request.SortDirection != nil {
		queryParams.Add("sortDirection", fmt.Sprintf("%v", request.SortDirection))
	}
	if request.Filter != nil {
		queryParams.Add("filter", fmt.Sprintf("%v", request.Filter))
	}
	if request.FilterField != nil {
		queryParams.Add("filterField", fmt.Sprintf("%v", request.FilterField))
	}
	if request.SearchValue != nil {
		queryParams.Add("searchValue", fmt.Sprintf("%v", request.SearchValue))
	}
	if request.SearchField != nil {
		queryParams.Add("searchField", fmt.Sprintf("%v", request.SearchField))
	}
	for _, value := range request.Ids {
		queryParams.Add("ids", fmt.Sprintf("%v", value))
	}
	if request.PageSize != nil {
		queryParams.Add("pageSize", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.PageNumber != nil {
		queryParams.Add("pageNumber", fmt.Sprintf("%v", *request.PageNumber))
	}
	if request.IncludeCounts != nil {
		queryParams.Add("includeCounts", fmt.Sprintf("%v", *request.IncludeCounts))
	}
	if request.IncludeLength != nil {
		queryParams.Add("includeLength", fmt.Sprintf("%v", *request.IncludeLength))
	}
	if request.IncludeLinks != nil {
		queryParams.Add("includeLinks", fmt.Sprintf("%v", *request.IncludeLinks))
	}
	if request.IncludeMessages != nil {
		queryParams.Add("includeMessages", fmt.Sprintf("%v", *request.IncludeMessages))
	}
	if request.For != nil {
		queryParams.Add("for", fmt.Sprintf("%v", request.For))
	}
	if request.Q != nil {
		queryParams.Add("q", fmt.Sprintf("%v", *request.Q))
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

	var response *flatfilego.GetRecordsResponse
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

// Updates existing records in a workbook sheet
//
// ID of sheet
func (c *Client) Update(ctx context.Context, sheetId flatfilego.SheetId, request flatfilego.Records) (*flatfilego.VersionResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/records", sheetId)

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

	var response *flatfilego.VersionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPut,
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

// Adds records to a workbook sheet
//
// ID of sheet
func (c *Client) Insert(ctx context.Context, sheetId flatfilego.SheetId, request []flatfilego.RecordData) (*flatfilego.RecordsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/records", sheetId)

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

	var response *flatfilego.RecordsResponse
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

// Deletes records from a workbook sheet
//
// ID of sheet
func (c *Client) Delete(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.DeleteRecordsRequest) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/records", sheetId)

	queryParams := make(url.Values)
	for _, value := range request.Ids {
		queryParams.Add("ids", fmt.Sprintf("%v", value))
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

// Searches for all values that match the 'find' value (globally or for a specific field via 'fieldKey') and replaces them with the 'replace' value. Wrap 'find' value in double quotes for exact match (""). Returns a commitId for the updated records
//
// ID of sheet
func (c *Client) FindAndReplace(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.FindAndReplaceRecordRequest) (*flatfilego.VersionResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/find-replace", sheetId)

	queryParams := make(url.Values)
	if request.Filter != nil {
		queryParams.Add("filter", fmt.Sprintf("%v", request.Filter))
	}
	if request.FilterField != nil {
		queryParams.Add("filterField", fmt.Sprintf("%v", request.FilterField))
	}
	if request.SearchValue != nil {
		queryParams.Add("searchValue", fmt.Sprintf("%v", request.SearchValue))
	}
	if request.SearchField != nil {
		queryParams.Add("searchField", fmt.Sprintf("%v", request.SearchField))
	}
	for _, value := range request.Ids {
		queryParams.Add("ids", fmt.Sprintf("%v", value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.VersionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPut,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
