// This file was auto-generated by Fern from our API Definition.

package sheets

import (
	bytes "bytes"
	context "context"
	flatfilego "github.com/FlatFilers/flatfile-go"
	core "github.com/FlatFilers/flatfile-go/core"
	internal "github.com/FlatFilers/flatfile-go/internal"
	option "github.com/FlatFilers/flatfile-go/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Returns sheets in a workbook
func (c *Client) List(
	ctx context.Context,
	request *flatfilego.ListSheetsRequest,
	opts ...option.RequestOption,
) (*flatfilego.ListSheetsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/sheets"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.ListSheetsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns a sheet in a workbook
func (c *Client) Get(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	opts ...option.RequestOption,
) (*flatfilego.SheetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v",
		sheetId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.SheetResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Deletes a specific sheet from a workbook
func (c *Client) Delete(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v",
		sheetId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Trigger data hooks and validation to run on a sheet
func (c *Client) Validate(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/validate",
		sheetId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns records from a sheet in a workbook as a csv file
func (c *Client) GetRecordsAsCsv(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	request *flatfilego.GetRecordsCsvRequest,
	opts ...option.RequestOption,
) (io.Reader, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/download",
		sheetId,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	response := bytes.NewBuffer(nil)
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns counts of records from a sheet
func (c *Client) GetRecordCounts(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	request *flatfilego.GetRecordCountsRequest,
	opts ...option.RequestOption,
) (*flatfilego.RecordCountsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/counts",
		sheetId,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.RecordCountsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns the commit versions for a sheet
func (c *Client) GetSheetCommits(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	request *flatfilego.ListSheetCommitsRequest,
	opts ...option.RequestOption,
) (*flatfilego.ListCommitsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/commits",
		sheetId,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.ListCommitsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Locks a sheet
func (c *Client) LockSheet(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/lock",
		sheetId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Removes a lock from a sheet
func (c *Client) UnlockSheet(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/unlock",
		sheetId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns record cell values grouped by all fields in the sheet
func (c *Client) GetCellValues(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	request *flatfilego.GetFieldValuesRequest,
	opts ...option.RequestOption,
) (*flatfilego.CellsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v/cells",
		sheetId,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.CellsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Updates Sheet
func (c *Client) UpdateSheet(
	ctx context.Context,
	// ID of sheet
	sheetId flatfilego.SheetId,
	request *flatfilego.SheetUpdateRequest,
	opts ...option.RequestOption,
) (*flatfilego.SheetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/sheets/%v",
		sheetId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.SheetResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
