// This file was auto-generated by Fern from our API Definition.

package accounts

import (
	context "context"
	flatfilego "github.com/FlatFilers/flatfile-go"
	core "github.com/FlatFilers/flatfile-go/core"
	option "github.com/FlatFilers/flatfile-go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Get the current account
func (c *Client) GetCurrent(
	ctx context.Context,
	opts ...option.RequestOption,
) (*flatfilego.AccountResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "accounts/current"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *flatfilego.AccountResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update the current account
func (c *Client) UpdateCurrent(
	ctx context.Context,
	request *flatfilego.AccountPatch,
	opts ...option.RequestOption,
) (*flatfilego.AccountResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "accounts/current"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *flatfilego.AccountResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPatch,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}