// This file was auto-generated by Fern from our API Definition.

package users

import (
	context "context"
	fmt "fmt"
	flatfilego "github.com/FlatFilers/flatfile-go"
	core "github.com/FlatFilers/flatfile-go/core"
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

// Gets a list of users
func (c *Client) List(ctx context.Context, request *flatfilego.ListUsersRequest) (*flatfilego.ListUsersResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "users"

	queryParams := make(url.Values)
	if request.Email != nil {
		queryParams.Add("email", fmt.Sprintf("%v", *request.Email))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.ListUsersResponse
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

// Gets a user
//
// The user id
func (c *Client) Get(ctx context.Context, userId flatfilego.UserId) (*flatfilego.UserResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v", userId)

	var response *flatfilego.UserResponse
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
