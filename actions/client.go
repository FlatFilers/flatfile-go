// This file was auto-generated by Fern from our API Definition.

package actions

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

func (c *Client) Create(
	ctx context.Context,
	request *flatfilego.Action,
	opts ...option.RequestOption,
) (flatfilego.ActionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/actions"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response flatfilego.ActionResponse
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

func (c *Client) BulkCreate(
	ctx context.Context,
	request flatfilego.Actions,
	opts ...option.RequestOption,
) (flatfilego.ActionsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := baseURL + "/actions/bulk"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response flatfilego.ActionsResponse
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

func (c *Client) Get(
	ctx context.Context,
	// The id of the action to return
	actionId flatfilego.ActionId,
	opts ...option.RequestOption,
) (flatfilego.ActionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/actions/%v",
		actionId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response flatfilego.ActionResponse
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

func (c *Client) Update(
	ctx context.Context,
	// The id of the action to patch
	actionId flatfilego.ActionId,
	request *flatfilego.ActionUpdate,
	opts ...option.RequestOption,
) (flatfilego.ActionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/actions/%v",
		actionId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response flatfilego.ActionResponse
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

func (c *Client) Delete(
	ctx context.Context,
	// The id of the action to delete
	actionId flatfilego.ActionId,
	opts ...option.RequestOption,
) (*flatfilego.Success, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.x.flatfile.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/actions/%v",
		actionId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

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
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
