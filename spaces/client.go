// This file was auto-generated by Fern from our API Definition.

package spaces

import (
	context "context"
	flatfilego "github.com/FlatFilers/flatfile-go"
	core "github.com/FlatFilers/flatfile-go/core"
	internal "github.com/FlatFilers/flatfile-go/internal"
	option "github.com/FlatFilers/flatfile-go/option"
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

// Returns all spaces for an account or environment
func (c *Client) List(
	ctx context.Context,
	request *flatfilego.ListSpacesRequest,
	opts ...option.RequestOption,
) (*flatfilego.ListSpacesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/spaces"
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
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.ListSpacesResponse
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Creates a new space based on an existing Space Config
func (c *Client) Create(
	ctx context.Context,
	request *flatfilego.SpaceConfig,
	opts ...option.RequestOption,
) (*flatfilego.SpaceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/spaces"
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

	var response *flatfilego.SpaceResponse
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
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns a single space
func (c *Client) Get(
	ctx context.Context,
	// ID of space to return
	spaceId flatfilego.SpaceId,
	opts ...option.RequestOption,
) (*flatfilego.SpaceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v",
		spaceId,
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

	var response *flatfilego.SpaceResponse
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete a space
func (c *Client) Delete(
	ctx context.Context,
	// ID of space to delete
	spaceId flatfilego.SpaceId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v",
		spaceId,
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

// Delete multiple spaces by id
func (c *Client) BulkDelete(
	ctx context.Context,
	request *flatfilego.DeleteSpacesRequest,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/spaces"
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

// Update a space, to change the name for example
func (c *Client) Update(
	ctx context.Context,
	// ID of space to update
	spaceId flatfilego.SpaceId,
	request *flatfilego.SpaceConfig,
	opts ...option.RequestOption,
) (*flatfilego.SpaceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v",
		spaceId,
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

	var response *flatfilego.SpaceResponse
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

// Sets archivedAt timestamp on a space
func (c *Client) ArchiveSpace(
	ctx context.Context,
	// ID of space to archive
	spaceId flatfilego.SpaceId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/archive",
		spaceId,
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

// Sets archivedAt timestamp on a space to null
func (c *Client) UnarchiveSpace(
	ctx context.Context,
	// ID of space to unarchive
	spaceId flatfilego.SpaceId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/unarchive",
		spaceId,
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

// Creates a new guidance
func (c *Client) CreateGuidance(
	ctx context.Context,
	// ID of the space
	spaceId flatfilego.SpaceId,
	request *flatfilego.GuidanceApiCreateData,
	opts ...option.RequestOption,
) (*flatfilego.GuidanceResource, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/guidance",
		spaceId,
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
		403: func(apiError *core.APIError) error {
			return &flatfilego.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.GuidanceResource
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
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Lists guidances
func (c *Client) ListGuidance(
	ctx context.Context,
	// ID of the space
	spaceId flatfilego.SpaceId,
	request *flatfilego.ListGuidanceRequest,
	opts ...option.RequestOption,
) (*flatfilego.GuidanceListResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/guidance",
		spaceId,
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
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &flatfilego.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.GuidanceListResponse
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieves a guidance by its id
func (c *Client) GetGuidance(
	ctx context.Context,
	// ID of the space
	spaceId flatfilego.SpaceId,
	// ID of the guidance
	guidanceId flatfilego.GuidanceId,
	request *flatfilego.GetGuidanceRequest,
	opts ...option.RequestOption,
) (*flatfilego.GuidanceResource, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/guidance/%v",
		spaceId,
		guidanceId,
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
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flatfilego.BadRequestError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &flatfilego.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.GuidanceResource
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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Updates a guidance with the given id
func (c *Client) UpdateGuidance(
	ctx context.Context,
	// ID of the space
	spaceId flatfilego.SpaceId,
	// ID of the guidance
	guidanceId flatfilego.GuidanceId,
	request *flatfilego.GuidanceApiUpdateData,
	opts ...option.RequestOption,
) (*flatfilego.GuidanceResource, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/guidance/%v",
		spaceId,
		guidanceId,
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
		403: func(apiError *core.APIError) error {
			return &flatfilego.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flatfilego.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *flatfilego.GuidanceResource
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

// Deletes a guidance by its id
func (c *Client) DeleteGuidance(
	ctx context.Context,
	// ID of the space
	spaceId flatfilego.SpaceId,
	// ID of the guidance
	guidanceId flatfilego.GuidanceId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/spaces/%v/guidance/%v",
		spaceId,
		guidanceId,
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
		403: func(apiError *core.APIError) error {
			return &flatfilego.ForbiddenError{
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
