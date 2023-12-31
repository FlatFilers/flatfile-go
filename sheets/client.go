// This file was auto-generated by Fern from our API Definition.

package sheets

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

// Returns sheets in a workbook
func (c *Client) List(ctx context.Context, request *flatfilego.ListSheetsRequest) (*flatfilego.ListSheetsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "sheets"

	queryParams := make(url.Values)
	queryParams.Add("workbookId", fmt.Sprintf("%v", request.WorkbookId))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.ListSheetsResponse
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

// Returns a sheet in a workbook
//
// ID of sheet
func (c *Client) Get(ctx context.Context, sheetId flatfilego.SheetId) (*flatfilego.SheetResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v", sheetId)

	var response *flatfilego.SheetResponse
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

// Deletes a specific sheet from a workbook
//
// ID of sheet
func (c *Client) Delete(ctx context.Context, sheetId flatfilego.SheetId) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v", sheetId)

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

// Trigger data hooks and validation to run on a sheet
//
// ID of sheet
func (c *Client) Validate(ctx context.Context, sheetId flatfilego.SheetId) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/validate", sheetId)

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
			Method:       http.MethodPost,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns records from a sheet in a workbook as a csv file
//
// ID of sheet
func (c *Client) GetRecordsAsCsv(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.GetRecordsCsvRequest) (io.Reader, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/download", sheetId)

	queryParams := make(url.Values)
	if request.VersionId != nil {
		queryParams.Add("versionId", fmt.Sprintf("%v", *request.VersionId))
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
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	response := bytes.NewBuffer(nil)
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns counts of records from a sheet
//
// ID of sheet
func (c *Client) GetRecordCounts(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.GetRecordCountsRequest) (*flatfilego.RecordCountsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/counts", sheetId)

	queryParams := make(url.Values)
	if request.VersionId != nil {
		queryParams.Add("versionId", fmt.Sprintf("%v", *request.VersionId))
	}
	if request.SinceVersionId != nil {
		queryParams.Add("sinceVersionId", fmt.Sprintf("%v", request.SinceVersionId))
	}
	if request.CommitId != nil {
		queryParams.Add("commitId", fmt.Sprintf("%v", request.CommitId))
	}
	if request.SinceCommitId != nil {
		queryParams.Add("sinceCommitId", fmt.Sprintf("%v", request.SinceCommitId))
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
	if request.ByField != nil {
		queryParams.Add("byField", fmt.Sprintf("%v", *request.ByField))
	}
	if request.Q != nil {
		queryParams.Add("q", fmt.Sprintf("%v", *request.Q))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.RecordCountsResponse
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

// Returns the commit versions for a sheet
//
// ID of sheet
func (c *Client) GetSheetCommits(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.ListSheetCommitsRequest) (*flatfilego.ListCommitsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/commits", sheetId)

	queryParams := make(url.Values)
	if request.Completed != nil {
		queryParams.Add("completed", fmt.Sprintf("%v", *request.Completed))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.ListCommitsResponse
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

// Locks a sheet
//
// ID of sheet
func (c *Client) LockSheet(ctx context.Context, sheetId flatfilego.SheetId) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/lock", sheetId)

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
			Method:       http.MethodPost,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Removes a lock from a sheet
//
// ID of sheet
func (c *Client) UnlockSheet(ctx context.Context, sheetId flatfilego.SheetId) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/unlock", sheetId)

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
			Method:       http.MethodPost,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns record cell values grouped by all fields in the sheet
//
// ID of sheet
func (c *Client) GetCellValues(ctx context.Context, sheetId flatfilego.SheetId, request *flatfilego.GetFieldValuesRequest) (*flatfilego.CellsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"sheets/%v/cells", sheetId)

	queryParams := make(url.Values)
	if request.FieldKey != nil {
		queryParams.Add("fieldKey", fmt.Sprintf("%v", request.FieldKey))
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
	if request.PageSize != nil {
		queryParams.Add("pageSize", fmt.Sprintf("%v", request.PageSize))
	}
	if request.PageNumber != nil {
		queryParams.Add("pageNumber", fmt.Sprintf("%v", request.PageNumber))
	}
	if request.Distinct != nil {
		queryParams.Add("distinct", fmt.Sprintf("%v", request.Distinct))
	}
	if request.IncludeCounts != nil {
		queryParams.Add("includeCounts", fmt.Sprintf("%v", request.IncludeCounts))
	}
	if request.SearchValue != nil {
		queryParams.Add("searchValue", fmt.Sprintf("%v", request.SearchValue))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.CellsResponse
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
