// This file was auto-generated by Fern from our API Definition.

package environments

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

// Get all environments
func (c *Client) List(
	ctx context.Context,
	request *flatfilego.ListEnvironmentsRequest,
	opts ...option.RequestOption,
) (*flatfilego.ListEnvironmentsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/environments"
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

	var response *flatfilego.ListEnvironmentsResponse
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

// Create a new environment
func (c *Client) Create(
	ctx context.Context,
	request *flatfilego.EnvironmentConfigCreate,
	opts ...option.RequestOption,
) (*flatfilego.EnvironmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/environments"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.EnvironmentResponse
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
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get a token which can be used to subscribe to events for this environment
func (c *Client) GetEnvironmentEventToken(
	ctx context.Context,
	request *flatfilego.GetEnvironmentEventTokenRequest,
	opts ...option.RequestOption,
) (*flatfilego.EventTokenResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/environments/subscription-token"
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

	var response *flatfilego.EventTokenResponse
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

// Returns a single environment
func (c *Client) Get(
	ctx context.Context,
	// ID of the environment to return. To fetch the current environment, pass `current`
	environmentId string,
	opts ...option.RequestOption,
) (*flatfilego.EnvironmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v",
		environmentId,
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

	var response *flatfilego.EnvironmentResponse
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

// Updates a single environment, to change the name for example
func (c *Client) Update(
	ctx context.Context,
	// ID of the environment to update
	environmentId string,
	request *flatfilego.EnvironmentConfigUpdate,
	opts ...option.RequestOption,
) (*flatfilego.Environment, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v",
		environmentId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *flatfilego.Environment
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
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Deletes a single environment
func (c *Client) Delete(
	ctx context.Context,
	// ID of the environment to delete
	environmentId string,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v",
		environmentId,
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

// Returns guides in an account
func (c *Client) ListGuides(
	ctx context.Context,
	// ID of the environment
	environmentId flatfilego.EnvironmentId,
	opts ...option.RequestOption,
) (*flatfilego.GuideListResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v/guides",
		environmentId,
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

	var response *flatfilego.GuideListResponse
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

// Returns a guide
func (c *Client) GetGuide(
	ctx context.Context,
	// ID of the environment
	environmentId flatfilego.EnvironmentId,
	// ID of guide
	guideId flatfilego.GuideId,
	opts ...option.RequestOption,
) (*flatfilego.GuideDetailResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v/guides/%v",
		environmentId,
		guideId,
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

	var response *flatfilego.GuideDetailResponse
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

// Updates a guide
func (c *Client) UpdateGuide(
	ctx context.Context,
	// ID of the environment
	environmentId flatfilego.EnvironmentId,
	// ID of guide
	guideId flatfilego.GuideId,
	request *flatfilego.GuideUpdateRequest,
	opts ...option.RequestOption,
) (*flatfilego.GuideDetailResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v/guides/%v",
		environmentId,
		guideId,
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

	var response *flatfilego.GuideDetailResponse
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

// Creates a guide
func (c *Client) CreateGuide(
	ctx context.Context,
	// ID of the environment
	environmentId flatfilego.EnvironmentId,
	request *flatfilego.GuideCreateRequest,
	opts ...option.RequestOption,
) (*flatfilego.GuideDetailResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v/guides",
		environmentId,
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

	var response *flatfilego.GuideDetailResponse
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

// Deletes a guide
func (c *Client) DeleteGuide(
	ctx context.Context,
	// ID of the environment
	environmentId flatfilego.EnvironmentId,
	// ID of guide to delete
	guideId flatfilego.GuideId,
	opts ...option.RequestOption,
) (*flatfilego.GuideDeleteResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v/guides/%v",
		environmentId,
		guideId,
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

	var response *flatfilego.GuideDeleteResponse
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

// Returns a specified version of a specific guide
func (c *Client) GetGuideVersion(
	ctx context.Context,
	// ID of the environment
	environmentId flatfilego.EnvironmentId,
	// ID of the guide
	guideId flatfilego.GuideId,
	// Version of the guide
	version int,
	opts ...option.RequestOption,
) (*flatfilego.GuideVersionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v/guides/%v/versions/%v",
		environmentId,
		guideId,
		version,
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

	var response *flatfilego.GuideVersionResponse
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
